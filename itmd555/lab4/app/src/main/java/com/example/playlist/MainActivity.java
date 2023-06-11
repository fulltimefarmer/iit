package com.example.playlist;

import android.annotation.SuppressLint;
import android.os.AsyncTask;
import android.os.Bundle;
import android.widget.Button;
import android.widget.ListView;

import androidx.appcompat.app.AlertDialog;
import androidx.appcompat.app.AppCompatActivity;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

import java.io.BufferedInputStream;
import java.io.BufferedReader;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;
import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;
import java.util.Objects;
import java.util.stream.Collectors;

public class MainActivity extends AppCompatActivity {

    private ListView playList;
    List<ItemObject> items;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        // Populate the data
        playList = findViewById(R.id.listView);
        new AsyncDataClass().execute();

        // Binding show dialog for button
        Button button = findViewById(R.id.analyseButton);
        button.setOnClickListener(v -> {
            AlertDialog.Builder builder = new AlertDialog.Builder(MainActivity.this);
            builder.setTitle("Data analytics");
            StringBuilder sb = new StringBuilder();
            sb.append("1.Total Number of CDs: ");
            sb.append(items.size());
            sb.append("\n");
            sb.append("2.Price of the most expensive CD: ");
            ItemObject item = items.stream().max(Comparator.comparing(ItemObject::getPrice)).get();
            sb.append(item.getPrice());
            sb.append("\n");
            sb.append("3.Countries of origin: ");
            sb.append(items.stream().map(ItemObject::getCountry).collect(Collectors.toList()));
            builder.setMessage(sb.toString());
            AlertDialog dialog = builder.create();
            dialog.show();
        });
    }

    @SuppressLint("StaticFieldLeak")
    private class AsyncDataClass extends AsyncTask<String, Void, String> {

        HttpURLConnection urlConnection;

        @Override
        protected String doInBackground(String... params) {
            StringBuilder jsonResult = new StringBuilder();
            try {
                URL url = new URL("http://www.papademas.net:81/cd_catalog.json");
                urlConnection = (HttpURLConnection) url.openConnection();
                InputStream in = new BufferedInputStream(urlConnection.getInputStream());
                BufferedReader reader = new BufferedReader(new InputStreamReader(in));
                String line;
                while ((line = reader.readLine()) != null) {
                    jsonResult.append(line);
                }
                System.out.println("Returned Json url object " + jsonResult.toString());
            } catch (Exception e) {
                System.out.println("Err: " + e);
            } finally {
                urlConnection.disconnect();
            }
            return jsonResult.toString();
        }

        @Override
        protected void onPreExecute() {}

        @Override
        protected void onPostExecute(String result) {
            System.out.println("Result on post execute: " + result);
            items = returnParsedJsonObject(result);
            items.sort(Comparator.comparing(ItemObject::getYear));
            CustomAdapter jsonCustomAdapter = new CustomAdapter(MainActivity.this, items);
            playList.setAdapter(jsonCustomAdapter);
        }
    }

    private List<ItemObject> returnParsedJsonObject(String result){
        List<ItemObject> jsonObject = new ArrayList<>();
        JSONObject resultObject;
        JSONArray jsonArray = null;
        ItemObject newItemObject;
        try {
            resultObject = new JSONObject(result);
            System.out.println("Preparsed JSON object " + resultObject.toString());
            // set up json Array to be parsed
            jsonArray = resultObject.optJSONArray("Bluesy_Music");
        } catch (JSONException e) {
            e.printStackTrace();
        }
        for(int i = 0; i < Objects.requireNonNull(jsonArray).length(); i++){
            JSONObject jsonChildNode;
            try {
                jsonChildNode = jsonArray.getJSONObject(i);
                //get all data from stream
                String sold = jsonChildNode.getString("SOLD");
                String title = jsonChildNode.getString("TITLE");
                String artist = jsonChildNode.getString("ARTIST");
                String country = jsonChildNode.getString("COUNTRY");
                String company = jsonChildNode.getString("COMPANY");
                String priceStr = jsonChildNode.getString("PRICE");
                Double price = Double.parseDouble(priceStr);
                String year = jsonChildNode.getString("YEAR");
                newItemObject = new ItemObject(sold, title, artist, country, company, price, year);
                jsonObject.add(newItemObject);
            } catch (JSONException e) {
                e.printStackTrace();
            }
        }
        return jsonObject;
    }

}