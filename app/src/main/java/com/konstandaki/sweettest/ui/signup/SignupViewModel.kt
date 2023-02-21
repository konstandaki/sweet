package com.konstandaki.sweettest.ui.signup

import android.util.Log
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.konstandaki.signup_service.SignupServiceGrpc
import com.konstandaki.signup_service.SignupServiceOuterClass
import com.konstandaki.sweettest.network.SignupClient
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch

class SignupViewModel() : ViewModel() {

    /*val homeUiState: StateFlow<HomeUiState> =
        itemsRepository.getAllItemsStream().map { HomeUiState(it) }
            .stateIn(
                scope = viewModelScope,
                started = SharingStarted.WhileSubscribed(TIMEOUT_MILLIS),
                initialValue = HomeUiState()
            )*/

    fun setPhone(phone: String) {
        viewModelScope.launch(context = Dispatchers.IO) {
            try {
                val signupClient = SignupClient("api.sweet.tv", 443)
                signupClient.setPhone(phone)
            } catch (e: Exception) {
                Log.d("CANDY", e.message!!)
            }
        }
    }
}

//data class HomeUiState(val itemList: List<Item> = listOf())