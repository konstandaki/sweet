package com.konstandaki.sweettest.data

import android.content.SharedPreferences
import com.konstandaki.sweettest.network.GrpcClient

const val AUTH_TOKEN_KEY = "AUTH_TOKEN"

interface SweetRepository {

    fun setPhone(phone: String): Int

    fun setCode(phone: String, code: Int): String

    fun saveAuthToken(token: String)

    fun getAuthToken(): String?

    fun close()
}

class GrpcSweetRepository(private val grpcClient: GrpcClient, private val sharedPrefs: SharedPreferences) : SweetRepository {

    override fun setPhone(phone: String): Int = grpcClient.setPhone(phone)

    override fun setCode(phone: String, code: Int): String = grpcClient.setCode(phone, code)

    // NOT SAFE
    override fun saveAuthToken(token: String) = sharedPrefs.edit().putString(AUTH_TOKEN_KEY, token).apply()

    override fun getAuthToken(): String? = sharedPrefs.getString(AUTH_TOKEN_KEY, "")

    override fun close() = grpcClient.close()
}