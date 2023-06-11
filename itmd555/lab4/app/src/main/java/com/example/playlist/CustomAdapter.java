package com.example.playlist;

import android.annotation.SuppressLint;
import android.content.Context;
import android.graphics.Color;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ArrayAdapter;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;

import java.util.List;

public class CustomAdapter extends ArrayAdapter<ItemObject> {

    private static class ViewHolder {
        TextView titleTextView;
        TextView yearTextView;
        TextView artistTextView;
    }

    public CustomAdapter(@NonNull Context context, List<ItemObject> objects) {
        super(context, 0, objects);
    }

    @SuppressLint({"SetTextI18n", "ResourceAsColor"})
    @NonNull
    @Override
    public View getView(int position, @Nullable View convertView, @NonNull ViewGroup parent) {
        ItemObject item = getItem(position);

        // Check if an existing view is being reused, otherwise inflate the view
        ViewHolder viewHolder = new ViewHolder();
        LayoutInflater inflater = LayoutInflater.from(getContext());
        convertView = inflater.inflate(R.layout.list, parent, false);
        viewHolder.titleTextView = convertView.findViewById(R.id.textView);
        viewHolder.yearTextView = convertView.findViewById(R.id.textView2);
        viewHolder.artistTextView = convertView.findViewById(R.id.textView3);
        convertView.setTag(viewHolder);

        if (position % 2 == 1) {
            convertView.setBackgroundColor(Color.LTGRAY);
        } else {
            convertView.setBackgroundColor(Color.WHITE);
        }

        // Populate the data into the template view using the data object
        viewHolder.titleTextView.setText("Song Title: " + item.getTitle());
        viewHolder.yearTextView.setText("Song Year: " + item.getYear());
        viewHolder.artistTextView.setText("Song Artist: " + item.getArtist());

        return convertView;
    }
}
