package com.konstandaki.sweettest.network

import android.util.Log
import androidx.compose.runtime.mutableStateOf
import com.konstandaki.signup_service.SignupServiceGrpc
import com.konstandaki.signup_service.SignupServiceOuterClass
import com.konstandaki.signup_service.SignupServiceOuterClass.SetPhoneRequest
import com.ua.mytrinity.tv_client.proto.Application
import com.ua.mytrinity.tv_client.proto.Application.ApplicationInfo
import com.ua.mytrinity.tv_client.proto.Application.ApplicationInfo.ApplicationType
import com.ua.mytrinity.tv_client.proto.Device
import com.ua.mytrinity.tv_client.proto.Device.DeviceInfo
import io.grpc.ManagedChannelBuilder
import io.grpc.stub.StreamObserver
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.asExecutor
import java.io.Closeable

class SignupClient(host: String, port: Int) : Closeable {

    val responseState = mutableStateOf("")

    private val channel = let {
        val builder = ManagedChannelBuilder.forAddress(host, port)
        builder.useTransportSecurity()
        builder.executor(Dispatchers.IO.asExecutor()).build()
    }

    private val stub = SignupServiceGrpc.newStub(channel)

    suspend fun setPhone(phone: String) {
        try {
            val request = SetPhoneRequest.newBuilder().setPhone(phone).setDevice(DeviceInfo
                .newBuilder().setType(DeviceInfo.DeviceType.DT_Android_Player)
                .setApplication(ApplicationInfo.newBuilder().setType(ApplicationType.AT_SWEET_TV_Player))
                .build()).build()
            //val request = SignupServiceGrpc.getSetPhoneMethod().toBuilder().build()
            stub.setPhone(request, object : StreamObserver<SignupServiceOuterClass.SetPhoneResponse> {
                override fun onNext(value: SignupServiceOuterClass.SetPhoneResponse?) {
                    Log.d("CANDY", "onNext")
                    Log.d("CANDY", value?.status?.name!!)
                }

                override fun onError(t: Throwable?) {
                    Log.d("CANDY", "onError")
                    Log.d("CANDY", t?.message.toString())
                }

                override fun onCompleted() {
                    Log.d("CANDY", "onCompleted")
                }
            })
        } catch (e: Exception) {
            Log.d("CANDY", e.message.toString())
        }
    }

    override fun close() {
        channel.shutdownNow()
    }
}