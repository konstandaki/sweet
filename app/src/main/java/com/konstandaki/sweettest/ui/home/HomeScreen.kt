package com.konstandaki.sweettest.ui.home

import androidx.compose.foundation.Image
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.grid.GridCells
import androidx.compose.foundation.lazy.grid.LazyVerticalGrid
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material.Button
import androidx.compose.material.Card
import androidx.compose.material.OutlinedTextField
import androidx.compose.material.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.input.KeyboardType
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import coil.compose.AsyncImage
import coil.request.ImageRequest
import com.konstandaki.sweettest.R

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
        is HomeUiState.Authorized -> PhotosGridScreen(modifier)
        is HomeUiState.Loading -> LoadingScreen(modifier)
        is HomeUiState.Success -> PhotosGridScreen(modifier)
        is HomeUiState.Error -> ErrorScreen(modifier)
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
        Text(stringResource(R.string.loading_failed))
        Button(onClick = {}/*retryAction*/) {
            Text(stringResource(R.string.OK))
        }
    }
}

/**
 * The home screen displaying photo grid.
 */
@Composable
fun PhotosGridScreen(/*photos: List<MarsPhoto>,*/ modifier: Modifier = Modifier) {
    LazyVerticalGrid(
        columns = GridCells.Adaptive(150.dp),
        modifier = modifier.fillMaxWidth(),
        contentPadding = PaddingValues(4.dp)
    ) {
        /*items(items = photos, key = { photo -> photo.id }) { photo ->
            MarsPhotoCard(photo)
        }*/
    }
}

@Composable
fun MarsPhotoCard(/*photo: MarsPhoto,*/ modifier: Modifier = Modifier) {
    Card(
        modifier = modifier
            .padding(4.dp)
            .fillMaxWidth()
            .aspectRatio(1f),
        elevation = 8.dp,
    ) {
        AsyncImage(
            model = ImageRequest.Builder(context = LocalContext.current)
                //.data(photo.imgSrc)
                .crossfade(true)
                .build(),
            contentDescription = stringResource(R.string.img),
            error = painterResource(R.drawable.ic_broken_image),
            placeholder = painterResource(R.drawable.loading_img),
            contentScale = ContentScale.FillBounds
        )
    }
}