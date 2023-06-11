package com.example.tempconverter;

import androidx.appcompat.app.AppCompatActivity;

import android.graphics.Color;
import android.os.Bundle;
import android.view.View;
import android.widget.EditText;
import android.widget.ImageView;
import android.widget.RadioButton;
import android.widget.Toast;

/**
 * Auth max zhou
 * Main cless for temp converter.
 */
public class MainActivity extends AppCompatActivity {

    private EditText text;

    private ImageView ivSun;

    private ImageView ivSnow;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        text = findViewById(R.id.editTextNumberDecimal);
    }

    /**
     * This method is called when user clicks the button and is handled because we assigned the name to the "OnClick property" of the button
     * */
    public void onClick(View view) {
        switch (view.getId()) {
            case R.id.button:
                RadioButton celsiusButton = findViewById(R.id.radioButton);
                RadioButton fahrenheitButton = findViewById(R.id.radioButton2);
                if (text.getText().length() == 0) {
                    Toast.makeText(this, "Please enter a valid number", Toast.LENGTH_LONG).show();
                    return;
                }
                float inputValue = Float.parseFloat(text.getText().toString());
                if (celsiusButton.isChecked()) {
                    text.setText(String.valueOf(ConverterUtil.convertCelsiusToFahrenheit(inputValue)));
                    celsiusButton.setChecked(false);
                    fahrenheitButton.setChecked(true);
                } else {
                    text.setText(String.valueOf(ConverterUtil.convertFahrenheitToCelsius(inputValue)));
                    fahrenheitButton.setChecked(false);
                    celsiusButton.setChecked(true);
                }
                // load and init images
                view = findViewById(R.id.activity_main);
                ivSun = findViewById(R.id.imageView);
                ivSun.setVisibility(View.INVISIBLE);
                ivSnow = findViewById(R.id.imageView1);
                ivSnow.setVisibility(View.INVISIBLE);
                if (inputValue > 90){
                    //set hex color to skyblue
                    view.setBackgroundColor(Color.parseColor("#87ceff"));
                    ivSnow.setVisibility(View.INVISIBLE);
                    ivSun.setVisibility(View.VISIBLE);
                    ivSun.setImageResource(R.drawable.sun);
                } else if (inputValue < 0) {
                    ivSun.setVisibility(View.INVISIBLE);
                    ivSnow.setVisibility(View.VISIBLE);
                    ivSnow.setVisibility(View.VISIBLE);
                    ivSnow.setImageResource(R.drawable.snow);
                } else {
                    view.setBackgroundColor(Color.YELLOW);
                    ivSun.setVisibility(View.INVISIBLE);
                    ivSnow.setVisibility(View.INVISIBLE);
                }
                break;
        }
    }

}