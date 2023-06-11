package com.example.playlist;

import androidx.appcompat.app.AppCompatActivity;

import android.content.Context;
import android.content.Intent;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;
import android.widget.Toast;

public class LoginActivity extends AppCompatActivity {

    private final String USERNAME = "user";
    private final String PASSWORD = "pwd";
    private final Integer MAXIMUM_RETRY_TIME = 3;
    private Integer CREDENTIAL_TRIED = 0;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_login);

        EditText username = findViewById(R.id.username);
        EditText password = findViewById(R.id.password);
        Button myButton = findViewById(R.id.loginButton);
        myButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Context that = getApplicationContext();
                if (CREDENTIAL_TRIED >= MAXIMUM_RETRY_TIME) {
                    Toast.makeText(that, "Exceed maximum retry time!", Toast.LENGTH_LONG).show();
                    return;
                }
                if (null != username.getText()
                    && null != password.getText()
                    && username.getText().toString().equals(USERNAME)
                    && password.getText().toString().equals(PASSWORD)) {
                    CREDENTIAL_TRIED = 0;
                    Toast.makeText(that, "Redirecting...", Toast.LENGTH_LONG).show();
                    Intent intent = new Intent(that, MainActivity.class);
                    startActivity(intent);
                } else {
                    CREDENTIAL_TRIED += 1;
                    Toast.makeText(that,
                            "Wrong Credentials, maximum retry time: " + (MAXIMUM_RETRY_TIME - CREDENTIAL_TRIED),
                            Toast.LENGTH_LONG).show();
                }
            }
        });
    }
}