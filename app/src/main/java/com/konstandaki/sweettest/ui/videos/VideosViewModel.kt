package com.konstandaki.sweettest.ui.videos

import androidx.lifecycle.ViewModel

class VideosViewModel() : ViewModel() {

    /*val homeUiState: StateFlow<HomeUiState> =
        itemsRepository.getAllItemsStream().map { HomeUiState(it) }
            .stateIn(
                scope = viewModelScope,
                started = SharingStarted.WhileSubscribed(TIMEOUT_MILLIS),
                initialValue = HomeUiState()
            )*/
}

//data class HomeUiState(val itemList: List<Item> = listOf())