package com.konstandaki.sweettest.ui.navigation

import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.navigation.NavHostController
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import com.konstandaki.sweettest.ui.player.PlayerDestination
import com.konstandaki.sweettest.ui.player.PlayerScreen
import com.konstandaki.sweettest.ui.signup.SignupDestination
import com.konstandaki.sweettest.ui.signup.SignupScreen

/**
 * Provides Navigation graph for the application.
 */
@Composable
fun SweetTestNavHost(
    navController: NavHostController,
    modifier: Modifier = Modifier,
) {
    NavHost(
        navController = navController,
        startDestination = SignupDestination.route,
        modifier = modifier
    ) {
        /*composable(route = PlayerDestination.route) {
            VideosScreen(
                navigateBack = { navController.popBackStack() },
                onNavigateUp = { navController.navigateUp() }
            )
        }*/
        composable(route = PlayerDestination.route) {
            PlayerScreen(
                //navigateBack = { navController.popBackStack() },
                //onNavigateUp = { navController.navigateUp() }
            )
        }
    }
}