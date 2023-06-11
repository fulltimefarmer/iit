package com.example.quiz;

import android.annotation.SuppressLint;
import android.os.AsyncTask;
import android.os.Bundle;
import android.os.CountDownTimer;
import android.view.View;
import android.widget.Button;
import android.widget.ImageView;
import android.widget.RadioButton;
import android.widget.RadioGroup;
import android.widget.RatingBar;
import android.widget.TextView;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;

import java.io.BufferedReader;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;
import java.util.Collections;
import java.util.LinkedList;
import java.util.List;
import java.util.Locale;

public class MainActivity extends AppCompatActivity {

    TextView txtView;
    TextView timerView;
    List<Question> questions = new LinkedList<>();

    static int questionNum = 0;
    static int starCount = 0;

    private RadioGroup radioQuestions;

    ImageView image;

    CountDownTimer timer = new CountDownTimer(600_000, 1000) {
        @SuppressLint("SetTextI18n")
        public void onTick(long millisUntilFinished) {
            long totalTime = 600000 - millisUntilFinished;
            long minutes = (totalTime / 1000) / 60;
            long seconds = (totalTime / 1000) % 60;
            String timeString = String.format(Locale.getDefault(), "%02d:%02d", minutes, seconds);
            timerView = findViewById(R.id.textViewTimer);
            timerView.setText("Time spend: " + timeString);
        }

        public void onFinish() {
            // Timer finished
        }
    };

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        BackgroundTask bt = new BackgroundTask();
        bt.execute("http://www.papademas.net:81/sample.txt"); //grab url
    }

    //background process to download the file from internet
    private class BackgroundTask extends AsyncTask<String, Integer, Void> {
        protected void onPreExecute() {
        }
        protected Void doInBackground(String... params) {
            URL url;
            String line;
            try {
                //create url object to point to the file location on internet
                url = new URL(params[0]);
                //make a request to server
                HttpURLConnection con = (HttpURLConnection) url.openConnection();
                //get InputStream instance
                InputStream is = con.getInputStream();
                //create BufferedReader object
                BufferedReader br = new BufferedReader(new InputStreamReader(is));
                int counter = 0;
                //read content of the file line by line & add it to Stringbuffer
                while ((line = br.readLine()) != null) {
                    String answer = counter == 4 ? "False" : "True";
                    Question q = new Question(line, answer);
                    questions.add(q);  //add to Arraylist
                    counter++;
                }
                br.close();
            } catch (Exception e) {
                e.printStackTrace();
            }
            return null;
        }
        protected void onPostExecute(Void result) {
            // Random the questions
            Collections.shuffle(questions);
            txtView = findViewById(R.id.textView1);
            //display read text in TextVeiw
            txtView.setText(questions.get(0).getStem());
            startQuiz();
        }
    }//end BackgroundTask class

    public void startQuiz() {
        timer.start();
        buttonListener();
        imageListener();
    }

    public void buttonListener() {
        Button btnDisplay = findViewById(R.id.btnDisplay);
        btnDisplay.setOnClickListener(v -> {
            if (checkResult(questionNum)) {
                Toast.makeText(MainActivity.this, " Right!", Toast.LENGTH_SHORT).show();
            } else {
                Toast.makeText(MainActivity.this, " Wrong!", Toast.LENGTH_SHORT).show();
            }
        });
    }//end buttonListener

    public void imageListener() {
        RatingBar rb = findViewById(R.id.ratingBar);
        image = findViewById(R.id.imageView1);
        image.setOnClickListener(view -> {
            if (checkResult(questionNum)) {
                starCount++;
            }
            if (questionNum == 4) {
                timer.cancel();
                rb.setRating(starCount);
                rb.setVisibility(View.VISIBLE);
            } else {
                txtView.setText(questions.get(++questionNum).getStem());
                radioQuestions.check(R.id.radioTrue);
            }
        });
    }//end imageListener

    private boolean checkResult(int questionNum) {
        // get selected radio button from radioGroup
        radioQuestions = findViewById(R.id.radioQuestions);
        int selectedId = radioQuestions.getCheckedRadioButtonId();
        // find the radiobutton by returned id
        RadioButton radioButton = findViewById(selectedId);
        return radioButton.getText().equals(questions.get(questionNum).getAnswer());
    }

    private static class Question {
        private final String stem;
        private final String answer;

        public Question(String stem, String answer) {
            this.stem = stem;
            this.answer = answer;
        }

        public String getStem() {
            return stem;
        }

        public String getAnswer() {
            return answer;
        }
    }
}//end activity