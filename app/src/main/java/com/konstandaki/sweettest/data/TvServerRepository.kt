package com.konstandaki.sweettest.data

import kotlinx.coroutines.flow.Flow

interface TvServerRepository {

    fun getAllChannelsStream(): Flow<List<Any>>
}