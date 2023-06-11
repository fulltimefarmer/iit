package com.example.tempconverter2;

import androidx.appcompat.app.AppCompatActivity;

import android.annotation.SuppressLint;
import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.view.ViewStub;
import android.widget.ArrayAdapter;
import android.widget.BaseAdapter;
import android.widget.Button;
import android.widget.CheckBox;
import android.widget.CompoundButton;
import android.widget.ImageView;
import android.widget.ListView;
import android.widget.SeekBar;
import android.widget.SeekBar.OnSeekBarChangeListener;
import android.widget.TextView;
import android.widget.Toast;


public class MainActivity extends AppCompatActivity {

    SeekBar seekBar;   //declare seekbar object
    TextView textView;
    TextView subTextView;
    Button button;
    //declare member variables for SeekBar
    int discrete = 0;
    int start = 50;
    int start_position = 50; //progress tracker
    int temp = 0;
    //declare objects for ViewStub
    ViewStub stub;
    CheckBox checkBox;
    ListView lv; //declare Listview object

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        //declare viewstub object
        stub = findViewById(R.id.viewStub1);
        @SuppressWarnings("unused")
        View inflated = stub.inflate();
        stub.setVisibility(View.INVISIBLE);

        button = findViewById(R.id.button);
        button.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                checkBox.setVisibility(View.VISIBLE);
                seekBar.setVisibility(View.VISIBLE);
                textView.setVisibility(View.VISIBLE);
                stub.setVisibility(View.GONE);
                checkBox.setChecked(false);
                seekBar.setProgress(22);
            }
        });
        button.setVisibility(View.INVISIBLE);

        subTextView = findViewById(R.id.subTextView);
        subTextView.setVisibility(View.INVISIBLE);

        //ViewStub logic
        checkBox = findViewById(R.id.checkBox1);

        //handle checkbox click event
        checkBox.setOnCheckedChangeListener(new CheckBox.OnCheckedChangeListener() {
            public void onCheckedChanged(CompoundButton arg0, boolean isChecked) {
                if (isChecked) {
                    //remove objs from parent view to allow for child view objs
                    checkBox.setVisibility(View.GONE);
                    seekBar.setVisibility(View.GONE);
                    textView.setVisibility(View.GONE);
                    stub.setVisibility(View.VISIBLE);
                    subTextView.setVisibility(View.VISIBLE);
                    button.setVisibility(View.VISIBLE);
                }
            }
        });

        //seekbar logic
        textView = findViewById(R.id.textview);
        textView.setText("	Fahrenheit at 32 degrees");

        //set default view
        seekBar = findViewById(R.id.seekbar);
        seekBar.setProgress(start_position);
        seekBar.incrementProgressBy(1);

        //create event handler for SeekBar
        seekBar.setOnSeekBarChangeListener(new OnSeekBarChangeListener() {
            @Override
            public void onStopTrackingTouch(SeekBar seekBar) {
                if (temp == 0) {  //for initial view result
                    Toast.makeText(getBaseContext(), "Fahrenheit result: " + "32 degrees", Toast.LENGTH_SHORT).show();
                } else {
                    Toast.makeText(getBaseContext(), "Fahrenheit result: " + discrete + " degrees", Toast.LENGTH_SHORT).show();
                }
            }

            @Override
            public void onStartTrackingTouch(SeekBar seekBar) {
            }

            @SuppressLint("SetTextI18n")
            @Override
            public void onProgressChanged(SeekBar seekBar, int progress, boolean fromUser) {
                // To convert progress passed as discrete (Fahrenheit) value
                temp = progress - start;
                discrete = (int) Math.round((((temp * 9.0) / 5.0) + 32)); //convert C to F temp
                textView.setText("	Fahrenheit at " + discrete + " degrees");
            }
        });

        //Listview logic
        String[] wkTemps = new String[]{
                "Monday\nHigh 27℃/Low 19℃",
                "Tuesday\n26℃/Low 22℃",
                "Wednesday\n26℃/Low 22℃",
                "Thursday\n23℃/Low 15℃",
                "Friday\n19℃/Low 15℃"};

        lv = findViewById(R.id.listView);
//        @SuppressWarnings({"unchecked", "rawtypes"})
        /*
         * To use a basic ArrayAdapter, you just need to initialize the adapter and
         * attach the adapter to the ListView. First, initialize the adapter...: *
         */
//                ArrayAdapter adapter = new ArrayAdapter(this, R.layout.list_item, R.id.content, wkTemps);
        // Assign adapter to ListView

        int[] imgs = {R.mipmap.img_sunny, R.mipmap.img_rain_snow,
                R.mipmap.img_sunny, R.mipmap.img_rain_snow, R.mipmap.img_rain};
        lv.setAdapter(new BaseAdapter() {
            @Override
            public int getCount() {
                return wkTemps.length;
            }

            @Override
            public String getItem(int i) {
                return wkTemps[i];
            }

            @Override
            public long getItemId(int i) {
                return i;
            }

            @Override
            public View getView(int position, View convertView, ViewGroup parent) {
                LayoutInflater inflater = LayoutInflater.from(MainActivity.this);
                View view = inflater.inflate(R.layout.list_item, null);
                TextView content = view.findViewById(R.id.content);
                content.setText(getItem(position));

                ImageView img = view.findViewById(R.id.img);
                img.setImageResource(imgs[position]);
                return view;
            }
        });


    }//end onCreate method

}