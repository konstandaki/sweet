package com.konstandaki.sweettest.data

import android.content.SharedPreferences
import com.konstandaki.sweettest.network.GrpcClient
import com.konstandaki.sweettest.proto.MovieServiceOuterClass
import tv_service.ChannelOuterClass

const val AUTH_TOKEN_KEY = "AUTH_TOKEN"

interface SweetRepository {

    fun setPhone(phone: String): Int

    fun setCode(phone: String, code: Int): String

    fun saveAuthToken(token: String)

    fun getAuthToken(): String?

    fun isAuthTokenPresent(): Boolean

    fun getChannels(): MutableList<ChannelOuterClass.Channel>?

    fun openStream(channelId: Int): GrpcClient.Stream?

    fun closeStream(streamId: Int): Int

    fun getGenres(): MutableList<MovieServiceOuterClass.Genre>?

    fun getGenreMovies(genreId: Int): MutableList<Int>?

    fun getMovieInfo(moviesIds: MutableList<Int>?): MutableList<MovieServiceOuterClass.Movie>?

    fun getMoviesByGenres(): MutableMap<MovieServiceOuterClass.Genre, MutableList<MovieServiceOuterClass.Movie>?>

    fun getLink(movieId: Int): String

    fun close()
}

class GrpcSweetRepository(private val grpcClient: GrpcClient, private val sharedPrefs: SharedPreferences) : SweetRepository {

    override fun setPhone(phone: String): Int = grpcClient.setPhone(phone)

    override fun setCode(phone: String, code: Int): String = grpcClient.setCode(phone, code)

    // NOT SAFE
    override fun saveAuthToken(token: String) = sharedPrefs.edit().putString(AUTH_TOKEN_KEY, token).apply()

    override fun getAuthToken(): String? = sharedPrefs.getString(AUTH_TOKEN_KEY, "")

    override fun isAuthTokenPresent(): Boolean = !getAuthToken().isNullOrBlank()

    override fun getChannels(): MutableList<ChannelOuterClass.Channel>? = grpcClient.getChannels(getAuthToken())

    override fun openStream(channelId: Int): GrpcClient.Stream? = grpcClient.openStream(getAuthToken(), channelId)

    override fun closeStream(streamId: Int): Int = grpcClient.closeStream(getAuthToken(), streamId)

    override fun getGenres(): MutableList<MovieServiceOuterClass.Genre>? = grpcClient.getGenres(getAuthToken())

    override fun getGenreMovies(genreId: Int): MutableList<Int>? = grpcClient.getGenreMovies(getAuthToken(), genreId)

    override fun getMovieInfo(moviesIds: MutableList<Int>?): MutableList<MovieServiceOuterClass.Movie>? = grpcClient.getMovieInfo(getAuthToken(), moviesIds)

    override fun getMoviesByGenres(): MutableMap<MovieServiceOuterClass.Genre, MutableList<MovieServiceOuterClass.Movie>?> = grpcClient.getMoviesByGenres(getAuthToken())

    override fun getLink(movieId: Int): String = grpcClient.getLink(getAuthToken(), movieId)

    override fun close() = grpcClient.close()
}