package com.konstandaki.sweettest.ui

import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.material.MaterialTheme
import androidx.compose.material.Scaffold
import androidx.compose.material.Surface
import androidx.compose.material.Text
import androidx.compose.material.TopAppBar
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.lifecycle.viewmodel.compose.viewModel
import com.konstandaki.sweettest.R
import com.konstandaki.sweettest.ui.signup.SignupScreen
import com.konstandaki.sweettest.ui.signup.SignupViewModel

@Composable
fun SweetTestApp(modifier: Modifier = Modifier) {
    Scaffold(
        modifier = modifier.fillMaxSize(),
        topBar = { TopAppBar(title = { Text(stringResource(R.string.app_name)) }) }
    ) {
        Surface(
            modifier = Modifier
                .fillMaxSize()
                .padding(it),
            color = MaterialTheme.colors.background
        ) {
            val signupViewModel: SignupViewModel = viewModel(factory = AppViewModelProvider.Factory)
            SignupScreen(
                onSignupClick = { signupViewModel.setPhone("380662856580") }
                //onSignupClick = { signupViewModel.setPhone("00000000000") }
                //marsUiState = marsViewModel.marsUiState,
                //retryAction = marsViewModel::getMarsPhotos
            )
        }
    }
}