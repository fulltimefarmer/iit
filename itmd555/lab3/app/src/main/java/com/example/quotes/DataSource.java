package com.example.quotes;

import java.util.ArrayList;

public class DataSource {

    private ArrayList<Integer> mPhotoPool;
    private ArrayList<Integer> mQuotePool;
    private ArrayList<Integer> mPhotoHdPool;

    public int getDataSourceLength() {
        return mPhotoPool.size();
    }

    public DataSource() {
        mPhotoPool = new ArrayList();
        mQuotePool = new ArrayList();
        mPhotoHdPool = new ArrayList();
        setupPhotoPool();
        setupQuotePool();
        setupPhotoHDPool();
    }

    public ArrayList<Integer> getmPhotoHdPool() {
        return mPhotoHdPool;
    }

    public ArrayList<Integer> getmPhotoPool() {
        return mPhotoPool;
    }

    public ArrayList<Integer> getmQuotePool() {
        return mQuotePool;
    }

    private void setupPhotoPool() {
        mPhotoPool.add(R.drawable.steve_1);
        mPhotoPool.add(R.drawable.steve_2);
        mPhotoPool.add(R.drawable.steve_3);
        mPhotoPool.add(R.drawable.steve_4);
        mPhotoPool.add(R.drawable.steve_5);
        mPhotoPool.add(R.drawable.musk_1);
        mPhotoPool.add(R.drawable.musk_2);
        mPhotoPool.add(R.drawable.musk_3);
        mPhotoPool.add(R.drawable.musk_4);
        mPhotoPool.add(R.drawable.musk_5);
    }
    private void setupQuotePool() {
        mQuotePool.add(R.string.quote_1);
        mQuotePool.add(R.string.quote_2);
        mQuotePool.add(R.string.quote_3);
        mQuotePool.add(R.string.quote_4);
        mQuotePool.add(R.string.quote_5);
        mQuotePool.add(R.string.quote_6);
        mQuotePool.add(R.string.quote_7);
        mQuotePool.add(R.string.quote_8);
        mQuotePool.add(R.string.quote_9);
        mQuotePool.add(R.string.quote_10);
    }
    private void setupPhotoHDPool() {
        mPhotoHdPool.add(R.drawable.steve_hd_1);
        mPhotoHdPool.add(R.drawable.steve_hd_2);
        mPhotoHdPool.add(R.drawable.steve_hd_3);
        mPhotoHdPool.add(R.drawable.steve_hd_4);
        mPhotoHdPool.add(R.drawable.steve_hd_5);
        mPhotoHdPool.add(R.drawable.musk_hd_1);
        mPhotoHdPool.add(R.drawable.musk_hd_2);
        mPhotoHdPool.add(R.drawable.musk_hd_3);
        mPhotoHdPool.add(R.drawable.musk_hd_4);
        mPhotoHdPool.add(R.drawable.musk_hd_5);
    }
}
