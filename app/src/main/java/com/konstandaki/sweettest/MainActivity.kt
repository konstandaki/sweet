package com.konstandaki.sweettest

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import com.konstandaki.sweettest.ui.SweetTestApp
import com.konstandaki.sweettest.ui.theme.SweetTestTheme

class MainActivity : ComponentActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        setContent {
            SweetTestTheme() {
                SweetTestApp()
            }
        }
    }

    override fun onDestroy() {
        super.onDestroy()
        val sweetRepository = (application as SweetTestApplication).container.sweetRepository
        sweetRepository.close()
    }
}