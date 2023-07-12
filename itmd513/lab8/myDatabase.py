import sqlite3
from contacts import *

def initTable():
    conn = sqlite3.connect('contacts.db')
    cursor = conn.cursor()
    cursor.execute('''
    CREATE TABLE IF NOT EXISTS max(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT, 
        phone TEXT)
    ''')
    for contact in contactlist:
        cursor.execute('INSERT INTO max (name, phone) VALUES (?, ?)', contact)
    conn.commit()
    cursor.close()
    conn.close()


def createContact(name, phone):
    conn = sqlite3.connect('contacts.db')
    cursor = conn.cursor()
    cursor.execute('INSERT INTO max (name, phone) VALUES (?, ?)', (name, phone))
    conn.commit()
    cursor.close()
    conn.close()


def readContact():
    conn = sqlite3.connect('contacts.db')
    cursor = conn.cursor()
    cursor.execute('SELECT id, name, phone FROM max')
    contacts = cursor.fetchall()
    cursor.close()
    conn.close()
    return contacts


def updateContact(id, new_name, new_phone):
    conn = sqlite3.connect('contacts.db')
    cursor = conn.cursor()
    cursor.execute('UPDATE max SET name = ?, phone = ? WHERE id = ?', (new_name, new_phone, id))
    conn.commit()
    cursor.close()
    conn.close()


def deleteContact(id):
    conn = sqlite3.connect('contacts.db')
    cursor = conn.cursor()
    cursor.execute('DELETE FROM max WHERE id = ?', (id,))
    conn.commit()
    cursor.close()
    conn.close()
