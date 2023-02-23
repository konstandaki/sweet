package com.konstandaki.sweettest.ui.home

import androidx.activity.compose.BackHandler
import androidx.compose.foundation.Image
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.grid.GridCells
import androidx.compose.foundation.lazy.grid.LazyVerticalGrid
import androidx.compose.foundation.lazy.grid.items
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalLifecycleOwner
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.input.KeyboardType
import androidx.compose.ui.unit.dp
import androidx.compose.ui.viewinterop.AndroidView
import androidx.lifecycle.Lifecycle
import androidx.lifecycle.LifecycleEventObserver
import androidx.lifecycle.viewmodel.compose.viewModel
import coil.compose.AsyncImage
import coil.request.ImageRequest
import com.google.android.exoplayer2.ExoPlayer
import com.google.android.exoplayer2.MediaItem
import com.google.android.exoplayer2.ui.StyledPlayerView
import com.google.android.exoplayer2.util.MimeTypes
import com.konstandaki.sweettest.R
import com.konstandaki.sweettest.network.GrpcClient
import com.konstandaki.sweettest.proto.MovieServiceOuterClass
import tv_service.ChannelOuterClass

@Composable
fun HomeScreen(
    modifier: Modifier = Modifier,
    homeViewModel: HomeViewModel = viewModel(
        factory = HomeViewModel.Factory
    )
) {
    when (homeViewModel.homeUiState) {
        is HomeUiState.NotAuthorized -> SignupScreen(
            signupUiState = homeViewModel.signupUiState,
            onPhoneChange = homeViewModel::updateSignupUiState,
            onCodeChange = homeViewModel::updateSignupUiState,
            onGetCodeClick = { homeViewModel.setPhone() },
            onSignupClick = { homeViewModel.setCode() }
        )
        is HomeUiState.Authorized -> homeViewModel.getData()
        is HomeUiState.Loading -> LoadingScreen(modifier)
        is HomeUiState.DataLoaded -> TabsScreen(homeViewModel.selectedTabIndex,
            (homeViewModel.homeUiState as HomeUiState.DataLoaded).channels,
            onChannelClick = { homeViewModel.openStream(it) },
            onUpdateTabIndex = { homeViewModel.updateSelectedTabIndex(it) })
        is HomeUiState.PlayStream -> PlayerScreen((homeViewModel.homeUiState as HomeUiState.PlayStream).stream,
            onBackPressed = { homeViewModel.closeStream(it) })
        is HomeUiState.Error -> ErrorScreen()
    }
}

/**
 * The home screen displaying the signup
 */
@Composable
fun SignupScreen(
    signupUiState: SignupUiState,
    onPhoneChange: (String, String) -> Unit,
    onCodeChange: (String, String) -> Unit,
    onSignupClick: () -> Unit,
    onGetCodeClick: () -> Unit,
    modifier: Modifier = Modifier
) {
    Column(
        modifier = modifier
            .padding(16.dp)
            .fillMaxWidth()
            .fillMaxHeight(),
        horizontalAlignment = Alignment.CenterHorizontally,
        verticalArrangement = Arrangement.spacedBy(8.dp)
    ) {
        OutlinedTextField(
            value = signupUiState.phone,
            onValueChange = { onPhoneChange(it, signupUiState.code) },
            keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Phone),
            label = { Text(stringResource(R.string.label_telephone)) },
            modifier = Modifier.fillMaxWidth(),
            singleLine = true
        )
        Spacer(modifier = Modifier.height(16.dp))
        OutlinedTextField(
            value = signupUiState.code,
            onValueChange = { onCodeChange(signupUiState.phone, it) },
            keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Number),
            label = { Text(stringResource(R.string.label_code)) },
            modifier = Modifier.fillMaxWidth(),
            singleLine = true
        )
        Spacer(modifier = Modifier.height(16.dp))
        Button(
            onClick = onGetCodeClick,
            enabled = signupUiState.isPhoneValid,
            modifier = Modifier.fillMaxWidth()
        ) {
            Text(stringResource(R.string.action_get_code))
        }
        Spacer(modifier = Modifier.height(16.dp))
        Button(
            onClick = onSignupClick,
            enabled = signupUiState.isAllValid,
            modifier = Modifier.fillMaxWidth()
        ) {
            Text(stringResource(R.string.action_signup))
        }
    }
}

/**
 * The home screen displaying the loading message.
 */
@Composable
fun LoadingScreen(modifier: Modifier = Modifier) {
    Box(
        contentAlignment = Alignment.Center,
        modifier = modifier.fillMaxSize()
    ) {
        Image(
            modifier = Modifier.size(200.dp),
            painter = painterResource(R.drawable.loading_img),
            contentDescription = stringResource(R.string.loading)
        )
    }
}

/**
 * The home screen displaying error message with re-attempt button.
 */
@Composable
fun ErrorScreen(modifier: Modifier = Modifier) {

    Column(
        modifier = modifier.fillMaxSize(),
        verticalArrangement = Arrangement.Center,
        horizontalAlignment = Alignment.CenterHorizontally
    ) {
        Text(stringResource(R.string.error))
    }
}

@Composable
fun TabsScreen(selectedTabIndex: Int, channels: MutableList<ChannelOuterClass.Channel>?,
               onChannelClick: (Int) -> Unit, onUpdateTabIndex: (Int) -> Unit) {
    val titles = listOf("CHANNELS", "MOVIES")
    Column {
        TabRow(selectedTabIndex = selectedTabIndex, divider = {}) {
            titles.forEachIndexed { index, title ->
                Tab(
                    text = { Text(title) },
                    selected = selectedTabIndex == index,
                    onClick = { onUpdateTabIndex(index) }
                )
            }
        }
        when (selectedTabIndex) {
            0 -> ChannelsGrid(channels = channels, onChannelClick = onChannelClick)
            //1 -> MoviesGrid(moviesByGenres = moviesByGenres, onMovieClick = onMovieClick)
        }
    }
}

@Composable
fun ChannelsGrid(channels: MutableList<ChannelOuterClass.Channel>?,
                     onChannelClick: (Int) -> Unit, modifier: Modifier = Modifier) {
    LazyVerticalGrid(
        columns = GridCells.Adaptive(150.dp),
        modifier = modifier.fillMaxWidth(),
        contentPadding = PaddingValues(4.dp),
    ) {
        items(items = channels?.toList()!!, key = { channel -> channel.id } ) { channel ->
            ChannelCard(channel, onChannelClick)
        }
    }
}

@Composable
fun ChannelCard(channel: ChannelOuterClass.Channel, onChannelClick: (Int) -> Unit, modifier: Modifier = Modifier) {
    Card(
        modifier = modifier
            .padding(4.dp)
            .fillMaxWidth()
            .clickable { onChannelClick(channel.id) }
            .aspectRatio(1.8f),
        elevation = 8.dp,
    ) {
        AsyncImage(
            model = ImageRequest.Builder(context = LocalContext.current)
                .data(channel.iconUrl)
                .crossfade(true)
                .build(),
            contentDescription = stringResource(R.string.img),
            error = painterResource(R.drawable.ic_broken_image),
            placeholder = painterResource(R.drawable.loading_img),
            contentScale = ContentScale.Inside
        )
    }
}

@Composable
fun MovieCard(movie: MovieServiceOuterClass.Movie, onMovieClick: (Int) -> Unit, modifier: Modifier = Modifier) {
    Card(
        modifier = modifier
            .fillMaxWidth()
            .clickable { onMovieClick(movie.id) }
            .aspectRatio(1f),
        elevation = 8.dp,
    ) {
        AsyncImage(
            model = ImageRequest.Builder(context = LocalContext.current)
                .data(movie.posterUrl)
                .crossfade(true)
                .build(),
            contentDescription = stringResource(R.string.img),
            error = painterResource(R.drawable.ic_broken_image),
            placeholder = painterResource(R.drawable.loading_img),
            contentScale = ContentScale.Fit
        )
    }
}

@Composable
fun PlayerScreen(stream: GrpcClient.Stream, onBackPressed: (Int) -> Unit) {
    val context = LocalContext.current

    val exoPlayer = ExoPlayer.Builder(LocalContext.current)
        .build()
        .also { exoPlayer ->
            val mediaItem = MediaItem.Builder()
                .setUri(stream.streamUrl)
                .setMimeType(MimeTypes.APPLICATION_M3U8)
                .build()
            exoPlayer.setMediaItem(mediaItem)
            exoPlayer.prepare()
        }

    val lifecycleOwner = rememberUpdatedState(LocalLifecycleOwner.current)

    BackHandler {
        onBackPressed(stream.streamId)
    }

    DisposableEffect(
        AndroidView(factory = {
            StyledPlayerView(context).apply {
                player = exoPlayer
            }
        })
    ) {
        val observer = LifecycleEventObserver { owner, event ->
            when (event) {
                Lifecycle.Event.ON_PAUSE -> {
                    exoPlayer.pause()
                }
                Lifecycle.Event.ON_RESUME -> {
                    exoPlayer.play()
                }
                else -> {}
            }
        }
        val lifecycle = lifecycleOwner.value.lifecycle
        lifecycle.addObserver(observer)

        onDispose {
            exoPlayer.release()
            lifecycle.removeObserver(observer)
        }
    }
}