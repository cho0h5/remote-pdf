package com.example.restremotecontroller

import android.os.Bundle
import android.os.StrictMode
import android.util.Log
import android.view.KeyEvent
import android.widget.TextView
import androidx.appcompat.app.AppCompatActivity
import java.net.URL


class MainActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        val policy =
            StrictMode.ThreadPolicy.Builder().permitAll().build()

        StrictMode.setThreadPolicy(policy)

        val status = URL("http://cho0h5.iptime.org:8080/status").readText()

        var statusText = findViewById<TextView>(R.id.status)
        statusText.setText(status)
    }

    override fun onKeyUp(keyCode: Int, event: KeyEvent): Boolean {
        return when (keyCode) {
            KeyEvent.KEYCODE_VOLUME_UP -> {
                Log.d("reponse:", URL("http://cho0h5.iptime.org:8080/prev").readText())
                true
            }
            KeyEvent.KEYCODE_VOLUME_DOWN -> {
                Log.d("reponse:", URL("http://cho0h5.iptime.org:8080/next").readText())
                true
            }
            else -> super.onKeyUp(keyCode, event)
        }
    }
}

