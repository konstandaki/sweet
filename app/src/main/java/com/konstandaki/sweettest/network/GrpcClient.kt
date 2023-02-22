package com.konstandaki.sweettest.network

import android.net.Uri
import com.konstandaki.signup_service.SignupServiceGrpc
import com.konstandaki.signup_service.SignupServiceOuterClass
import com.konstandaki.signup_service.SignupServiceOuterClass.SetPhoneRequest
import com.ua.mytrinity.tv_client.proto.Application.ApplicationInfo
import com.ua.mytrinity.tv_client.proto.Application.ApplicationInfo.ApplicationType
import com.ua.mytrinity.tv_client.proto.Device.DeviceInfo
import io.grpc.ManagedChannelBuilder
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.asExecutor
import java.io.Closeable

class GrpcClient(uri: Uri) : Closeable {

    companion object {
        const val SUCCESS_CODE = 0
        const val FAIL_CODE = -1
        const val ERROR = "ERROR"
    }

    private val channel = let {
        val builder = ManagedChannelBuilder.forAddress(uri.host, uri.port)
        builder.useTransportSecurity()
        builder.executor(Dispatchers.IO.asExecutor()).build()
    }

    private val signupStub = SignupServiceGrpc.newBlockingStub(channel)
    //private val tvStub = TvServiceGrpc.newStub(channel)

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

    override fun close() {
        channel.shutdownNow()
    }
}