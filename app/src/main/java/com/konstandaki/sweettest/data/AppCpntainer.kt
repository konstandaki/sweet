package com.konstandaki.sweettest.data

import android.content.Context

/**
 * App container for Dependency injection.
 */
interface AppContainer {
    val tvServerRepository: TvServerRepository
    val movieRepository: MovieRepository
}

/*class AppDataContainer(private val context: Context) : AppContainer {

    /*override val movieRepository: MovieRepository by lazy {
        OfflineItemsRepository(InventoryDatabase.getDatabase(context).itemDao())
    }*/
}*/