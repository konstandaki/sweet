package com.konstandaki.sweettest

import android.app.Application
import com.konstandaki.sweettest.data.AppContainer
import com.konstandaki.sweettest.data.AppDataContainer

class SweetTestApplication : Application() {

    /**
     * AppContainer instance used by the rest of classes to obtain dependencies
     */
    lateinit var container: AppContainer
    override fun onCreate() {
        super.onCreate()
        container = AppDataContainer(this)
    }
}