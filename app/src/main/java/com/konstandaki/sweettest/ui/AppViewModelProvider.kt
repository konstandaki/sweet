package com.konstandaki.sweettest.ui

import androidx.lifecycle.ViewModelProvider.AndroidViewModelFactory
import androidx.lifecycle.viewmodel.CreationExtras
import androidx.lifecycle.viewmodel.initializer
import androidx.lifecycle.viewmodel.viewModelFactory
import com.konstandaki.sweettest.SweetTestApplication
import com.konstandaki.sweettest.ui.player.PlayerViewModel
import com.konstandaki.sweettest.ui.signup.SignupViewModel
import com.konstandaki.sweettest.ui.videos.VideosViewModel

object AppViewModelProvider {
    val Factory = viewModelFactory {
        // Initializer for ItemEditViewModel
        initializer {
            SignupViewModel(
                // to do
            )
        }
        // Initializer for ItemEntryViewModel
        initializer {
            VideosViewModel(
                // to do
            )
        }

        // Initializer for ItemDetailsViewModel
        initializer {
            PlayerViewModel(
                // to do
            )
        }
    }
}

fun CreationExtras.sweetTestApplication(): SweetTestApplication =
    (this[AndroidViewModelFactory.APPLICATION_KEY] as SweetTestApplication)