package com.konstandaki.sweettest.network

import android.net.Uri
import com.konstandaki.signup_service.SignupServiceGrpc
import com.konstandaki.signup_service.SignupServiceOuterClass
import com.konstandaki.signup_service.SignupServiceOuterClass.SetPhoneRequest
import com.konstandaki.sweettest.proto.MovieServiceGrpc
import com.konstandaki.sweettest.proto.MovieServiceOuterClass
import com.konstandaki.sweettest.proto.MovieServiceOuterClass.GetConfigurationRequest
import com.konstandaki.sweettest.proto.MovieServiceOuterClass.GetGenreMoviesRequest
import com.konstandaki.sweettest.proto.MovieServiceOuterClass.GetMovieInfoRequest
import com.ua.mytrinity.tv_client.proto.Application.ApplicationInfo
import com.ua.mytrinity.tv_client.proto.Application.ApplicationInfo.ApplicationType
import com.ua.mytrinity.tv_client.proto.Device.DeviceInfo
import io.grpc.ManagedChannelBuilder
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.asExecutor
import tv_service.ChannelOuterClass
import tv_service.TvServiceGrpc
import tv_service.TvServiceOuterClass
import tv_service.TvServiceOuterClass.CloseStreamRequest
import java.io.Closeable

class GrpcClient(uri: Uri) : Closeable {

    companion object {
        const val SUCCESS_CODE = 0
        const val FAIL_CODE = -1
        const val ERROR = "ERROR"
    }

    data class Stream(val streamId: Int, val streamUrl: String)

    private val channel = let {
        val builder = ManagedChannelBuilder.forAddress(uri.host, uri.port)
        builder.useTransportSecurity()
        builder.executor(Dispatchers.IO.asExecutor()).build()
    }

    private val signupStub = SignupServiceGrpc.newBlockingStub(channel)
    private val tvStub = TvServiceGrpc.newBlockingStub(channel)
    private val movieStub = MovieServiceGrpc.newBlockingStub(channel)

    fun setPhone(phone: String): Int {
        return try {
            val request = SetPhoneRequest.newBuilder()
                .setPhone(phone)
                .setDevice(DeviceInfo.newBuilder()
                    .setType(DeviceInfo.DeviceType.DT_Android_Player)
                    .setApplication(ApplicationInfo.newBuilder()
                        .setType(ApplicationType.AT_SWEET_TV_Player)).build())
                .build()
            val response = signupStub.setPhone(request)
            response.status.number
        } catch (e: Exception) {
            FAIL_CODE
        }
    }

    fun setCode(phone: String, code: Int): String {
        return try {
            val request = SignupServiceOuterClass.SetCodeRequest.newBuilder().setPhone(phone).setConfirmationCode(code).build()
            val response = signupStub.setCode(request)
            if (response.status.number == SUCCESS_CODE) {
                response.authToken
            } else {
                ERROR
            }
        } catch (e: Exception) {
            ERROR
        }
    }

    fun getChannels(auth: String?): MutableList<ChannelOuterClass.Channel>? {
        try {
            val request = TvServiceOuterClass.GetChannelsRequest.newBuilder()
                .setAuth(auth)
                .setNeedIcons(true)
                .setNeedEpg(false)
                .setNeedOffsets(false)
                .build()
            val response = tvStub.getChannels(request)
            return if (response.status.number == SUCCESS_CODE && response.listList != null && response.listList.isNotEmpty()) {
                response.listList
            } else {
                null
            }
        } catch (e: Exception) {
            return null
        }
    }

    fun openStream(auth: String?, channelId: Int): Stream? {
        return try {
            val request = TvServiceOuterClass.OpenStreamRequest.newBuilder()
                .setAuth(auth)
                .setChannelId(channelId)
                .build()
            val response = tvStub.openStream(request)
            if (response.result.number == SUCCESS_CODE && response.hasHttpStream()) {
                Stream(response.streamId, buildStreamUrl(response.httpStream.host.address, response.httpStream.url))
            } else {
                null
            }
        } catch (e: Exception) {
            null
        }
    }

    private fun buildStreamUrl(host: String, url: String): String {
        return "http://$host$url"
    }

    fun closeStream(auth: String?, streamId: Int): Int {
        return try {
            val request = CloseStreamRequest.newBuilder()
                .setAuth(auth)
                .setStreamId(streamId)
                .build()
            val response = tvStub.closeStream(request)
            response.result.number
        } catch (e: Exception) {
            FAIL_CODE
        }
    }

    fun getGenres(auth: String?): MutableList<MovieServiceOuterClass.Genre>? {
        return try {
            val request = GetConfigurationRequest.newBuilder()
                .setAuth(auth)
                .build()
            val response = movieStub.getConfiguration(request)
            if (response.result.number == SUCCESS_CODE && response.genresList != null && response.genresList.isNotEmpty()) {
                response.genresList
            } else {
                null
            }
        } catch (e: Exception) {
            null
        }
    }

    fun getGenreMovies(auth: String?, genreId: Int): MutableList<Int>? {
        return try {
            val request = GetGenreMoviesRequest.newBuilder()
                .setAuth(auth)
                .setGenreId(genreId)
                .build()
            val response = movieStub.getGenreMovies(request)
            if (response.result.number == SUCCESS_CODE && response.moviesList != null && response.moviesList.isNotEmpty()) {
                response.moviesList
            } else {
                null
            }
        } catch (e: Exception) {
            null
        }
    }

    fun getMovieInfo(auth: String?, moviesIds: MutableList<Int>?): MutableList<MovieServiceOuterClass.Movie>? {
        return try {
            val request = GetMovieInfoRequest.newBuilder()
                .setAuth(auth)
                .addAllMovies(moviesIds)
                .build()
            val response = movieStub.getMovieInfo(request)
            if (response.result.number == SUCCESS_CODE && response.moviesList != null && response.moviesList.isNotEmpty()) {
                response.moviesList
            } else {
                null
            }
        } catch (e: Exception) {
            null
        }
    }

    fun getMoviesByGenres(auth: String?): MutableMap<MovieServiceOuterClass.Genre, MutableList<MovieServiceOuterClass.Movie>?> {
        val moviesByGenres = mutableMapOf<MovieServiceOuterClass.Genre, MutableList<MovieServiceOuterClass.Movie>?>()
        val genres = getGenres(auth)
        genres?.forEach {
            val moviesIds = getGenreMovies(auth, it.id)
            val movies = getMovieInfo(auth, moviesIds)
            moviesByGenres[it] = movies
        }
        return moviesByGenres
    }

    fun getLink(auth: String?, movieId: Int): String {
        return try {
            val request = MovieServiceOuterClass.GetLinkRequest.newBuilder()
                .setAuth(auth)
                .setMovieId(movieId)
                .build()
            val response = movieStub.getLink(request)
            if (response.status.number == SUCCESS_CODE && response.hasUrl()) {
                response.url
            } else {
                ERROR
            }
        } catch (e: Exception) {
            ERROR
        }
    }

    override fun close() {
        channel.shutdownNow()
    }
}