package com.konstandaki.sweettest

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import com.konstandaki.sweettest.ui.SweetTestApp
import com.konstandaki.sweettest.ui.theme.SweetTestTheme

/**
 * Activity for cupcake order flow.
 */
class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            SweetTestTheme() {
                SweetTestApp()
            }
        }
    }
}