package com.konstandaki.sweettest.data

import android.content.Context
import android.content.SharedPreferences
import android.net.Uri
import com.konstandaki.sweettest.network.GrpcClient

/**
 * App container for Dependency injection.
 */
interface AppContainer {
    val sharedPrefs: SharedPreferences
    val sweetRepository: SweetRepository
}

const val SERVER_URL = "https://api.sweet.tv:443/"
const val SHARED_PREFS_KEY = "auth_token"

class AppDataContainer(private val context: Context) : AppContainer {

    private val uri by lazy { Uri.parse(SERVER_URL) }
    private val grpcClient by lazy { GrpcClient(uri) }

    override val sharedPrefs: SharedPreferences = context.getSharedPreferences(SHARED_PREFS_KEY, Context.MODE_PRIVATE)

    override val sweetRepository: SweetRepository by lazy {
        GrpcSweetRepository(grpcClient, sharedPrefs)
    }
}