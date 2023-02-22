package com.konstandaki.sweettest.ui.home

import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.viewModelScope
import androidx.lifecycle.viewmodel.initializer
import androidx.lifecycle.viewmodel.viewModelFactory
import com.konstandaki.sweettest.SweetTestApplication
import com.konstandaki.sweettest.data.SweetRepository
import com.konstandaki.sweettest.network.GrpcClient.Companion.ERROR
import com.konstandaki.sweettest.network.GrpcClient.Companion.SUCCESS_CODE
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch

/**
 * UI state for the Home screen
 */
sealed interface HomeUiState {
    data class Success(val photos: List<Any>) : HomeUiState
    object Error : HomeUiState
    object Loading : HomeUiState
    object NotAuthorized : HomeUiState
    object Authorized : HomeUiState
}

data class SignupUiState(
    val phone: String = "",
    val code: String = "",
    val isPhoneValid: Boolean = false,
    val isAllValid: Boolean = false
)

class HomeViewModel(private val sweetRepository: SweetRepository) : ViewModel() {

    var homeUiState: HomeUiState by mutableStateOf(HomeUiState.NotAuthorized)
        private set

    var signupUiState by mutableStateOf(SignupUiState())
        private set

    fun updateSignupUiState(phone: String, code: String) {
        signupUiState = SignupUiState(phone, code, validatePhone(phone), validateAll(phone, code))
    }

    private fun validatePhone(phone: String): Boolean {
       return android.util.Patterns.PHONE.matcher(phone).matches()
    }

    private fun validateAll(phone: String, code: String): Boolean {
        return android.util.Patterns.PHONE.matcher(phone).matches() && (code.isNotBlank() && code.length == 4)
    }

    fun setPhone() {
        if (signupUiState.isPhoneValid) {
            homeUiState = HomeUiState.Loading
            viewModelScope.launch(context = Dispatchers.IO) {
                val result = sweetRepository.setPhone(signupUiState.phone.removePrefix("+"))
                homeUiState = if (result == SUCCESS_CODE) {
                    HomeUiState.NotAuthorized
                } else {
                    HomeUiState.Error
                }
            }
        }
    }

    fun setCode() {
        if (signupUiState.isAllValid) {
            homeUiState = HomeUiState.Loading
            viewModelScope.launch(context = Dispatchers.IO) {
                val result = sweetRepository.setCode(signupUiState.phone, signupUiState.code.toInt())
                homeUiState = if (result == ERROR) {
                    HomeUiState.Error
                } else {
                    sweetRepository.saveAuthToken(result)
                    HomeUiState.Authorized
                }
            }
        }
    }

    companion object {
        val Factory: ViewModelProvider.Factory = viewModelFactory {
            initializer {
                val application = (this[ViewModelProvider.AndroidViewModelFactory.APPLICATION_KEY] as SweetTestApplication)
                val sweetRepository = application.container.sweetRepository
                HomeViewModel(sweetRepository = sweetRepository)
            }
        }
    }
}