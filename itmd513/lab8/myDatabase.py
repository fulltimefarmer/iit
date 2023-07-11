import sqlite3
from contacts import *

def initTable():
    conn = sqlite3.connect('contacts.db')
    cursor = conn.cursor()
    cursor.execute('DROP TABLE IF EXISTS contacts')
    cursor.execute('''
    CREATE TABLE IF NOT EXISTS contacts(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT, 
        phone TEXT)
    ''')
    for contact in contactlist:
        cursor.execute('INSERT INTO contacts (name, phone) VALUES (?, ?)', contact)
    conn.commit()
    cursor.close()
    conn.close()


def createContact(name, phone):
    """
    Create a new contact in the database with the given name and phone number.
    Args:
    - name (str): The name of the contact.
    - phone (str): The phone number of the contact.
    """
    conn = sqlite3.connect('contacts.db')
    cursor = conn.cursor()
    cursor.execute('INSERT INTO contacts (name, phone) VALUES (?, ?)', (name, phone))
    conn.commit()
    cursor.close()
    conn.close()


def readContact():
    """
    Read all the contacts from the database.
    Returns:
    - list: A list of tuples containing the id, name, and phone number of each contact.
    """
    conn = sqlite3.connect('contacts.db')
    cursor = conn.cursor()
    cursor.execute('SELECT id, name, phone FROM contacts')
    contacts = cursor.fetchall()
    cursor.close()
    conn.close()
    return contacts


def updateContact(id, new_name, new_phone):
    """
    Update the name and phone number of a contact with the given id.
    Args:
    - id (int): The id of the contact to be updated.
    - new_name (str): The new name of the contact.
    - new_phone (str): The new phone number of the contact.
    """
    conn = sqlite3.connect('contacts.db')
    cursor = conn.cursor()
    cursor.execute('UPDATE contacts SET name = ?, phone = ? WHERE id = ?', (new_name, new_phone, id))
    conn.commit()
    cursor.close()
    conn.close()


def deleteContact(id):
    """
    Delete a contact with the given id from the database.
    Args:
    - id (int): The id of the contact to be deleted.
    """
    conn = sqlite3.connect('contacts.db')
    cursor = conn.cursor()
    cursor.execute('DELETE FROM contacts WHERE id = ?', (id,))
    conn.commit()
    cursor.close()
    conn.close()
