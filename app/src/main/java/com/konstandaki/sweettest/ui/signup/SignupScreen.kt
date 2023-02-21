package com.konstandaki.sweettest.ui.signup

import androidx.compose.foundation.layout.*
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material.Button
import androidx.compose.material.OutlinedTextField
import androidx.compose.material.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.input.KeyboardType
import androidx.compose.ui.unit.dp
import com.konstandaki.sweettest.R
import com.konstandaki.sweettest.ui.navigation.NavigationDestination

object SignupDestination : NavigationDestination {
    override val route = "signup"
    override val titleRes = R.string.screen_signup
}

@Composable
fun SignupScreen(
    //itemUiState: ItemUiState,
    //onItemValueChange: (ItemDetails) -> Unit,
    onSignupClick: () -> Unit,
    modifier: Modifier = Modifier
) {
    Column(
        modifier = modifier.padding(16.dp).fillMaxWidth(),
        horizontalAlignment = Alignment.CenterHorizontally,
        verticalArrangement = Arrangement.spacedBy(8.dp)
    ) {
        OutlinedTextField(
            value = /*itemDetails.name*/"",
            onValueChange = { /*onValueChange(itemDetails.copy(name = it))*/ },
            label = { Text(stringResource(R.string.label_telephone)) },
            modifier = Modifier.fillMaxWidth(),
            //enabled = enabled,
            singleLine = true
        )
        Spacer(modifier = Modifier.height(16.dp))
        OutlinedTextField(
            value = /*itemDetails.price*/"",
            onValueChange = { /*onValueChange(itemDetails.copy(price = it))*/ },
            keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Decimal),
            label = { Text(stringResource(R.string.label_code)) },
            modifier = Modifier.fillMaxWidth(),
            //enabled = enabled,
            singleLine = true
        )
        Spacer(modifier = Modifier.height(16.dp))
        Button(
            onClick = onSignupClick,
            //enabled = itemUiState.isEntryValid,
            modifier = Modifier.fillMaxWidth()
        ) {
            Text(stringResource(R.string.action_signup))
        }
    }
}