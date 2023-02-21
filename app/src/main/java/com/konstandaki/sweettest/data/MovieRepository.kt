package com.konstandaki.sweettest.data

import kotlinx.coroutines.flow.Flow

interface MovieRepository {

    fun getGenreMoviesStream(): Flow<List<Any>>
}