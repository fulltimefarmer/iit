import os
from tkinter import *
from tkinter import messagebox
from myDatabase import *

# Jun Zhou
# July 12th, 2023
# Lab8
# ITMD513

def selection():
    try:
        index = int(select.curselection()[0])
        current = contacts[index]
        global currentId 
        currentId = current[0]
        print('Current selected record: %d %s %s' % (current))
        return current
    except:
        messagebox.showerror("Error", "None of record is selected!")


def addContact():
    try:
        printList()
        createContact(nameVar.get(), phoneVar.get())
    except e:
        print(e)
        messagebox.showwarning("Warning", "Please enter a value in both the fields.")
    else:
        setList()
        messagebox.showinfo("Success", "Contacts added successfully!")


def editContact():
    try:
        printList()
        updateContact(currentId, nameVar.get(), phoneVar.get())
    except:
        messagebox.showwarning("Warning", "Please enter valid value in both the fields.")
    else:
        setList()
        messagebox.showinfo("Success", "Contacts added successfully!")


def removeContact():
    _, name, _ = selection()
    result = messagebox.askquestion('Confirmation', ('Are you sure you want to delete %s?' % name))
    if result == 'yes':
        printList()
        deleteContact(currentId)
        setList()


def loadContact():
    _, name, phone = selection()
    nameVar.set(name)
    phoneVar.set(phone)


def exitProgram():
    result = messagebox.askquestion('Exit', 'Are you want to exit?')
    if result == 'yes':
        os._exit(1)


def buildFrame():
    global nameVar, phoneVar, select
    root = Tk()
    root.title('My Contact List')
    frame1 = Frame(root)
    frame1.pack()
    Label(frame1, text='Name:').grid(row=0, column=0, sticky=W)
    nameVar = StringVar()
    name = Entry(frame1, textvariable=nameVar)
    name.grid(row=0, column=1, sticky=W)
    Label(frame1, text='Phone:').grid(row=1, column=0, sticky=W)
    phoneVar= StringVar()
    phone = Entry(frame1, textvariable=phoneVar)
    phone.grid(row=1, column=1, sticky=W)
    frame1 = Frame(root) # add a row of buttons
    frame1.pack()
    btn1 = Button(frame1,text=' Add  ',command=addContact)
    btn2 = Button(frame1,text='Update',command=editContact)
    btn3 = Button(frame1,text='Delete',command=removeContact)
    btn4 = Button(frame1,text=' Load ',command=loadContact)
    btn1.pack(side=LEFT)
    btn2.pack(side=LEFT)
    btn3.pack(side=LEFT)
    btn4.pack(side=LEFT)
    frame1 = Frame(root) # allow for selection of names
    frame1.pack()
    scroll = Scrollbar(frame1, orient=VERTICAL)
    select = Listbox(frame1, yscrollcommand=scroll.set, height=8)
    scroll.config (command=select.yview)
    scroll.pack(side=RIGHT, fill=Y)
    select.pack(side=LEFT,  fill=BOTH)
    frame1 = Frame(root) # add exit button
    frame1.pack()
    btn6 = Button(frame1, text=' Exit ', command=exitProgram)
    btn6.pack()
    return root


def printList():
    print('Current contacts:')
    for record in contacts:
        print(record)


def setList():
    global contacts
    contacts = readContact()
    contacts.sort(key = lambda x : x[1])
    print('Contacts after opertion:')
    for record in contacts:
        print(record)
    select.delete(0, END)
    for _, name, _ in contacts:
        select.insert(END, name)


if __name__ == "__main__":
    # the main program
    # initialize the application by building the GUI elements
    root = buildFrame()
    # initialize the database table / load records
    initTable()
    # set the contents of the list initially
    setList()
    # to keep the program from exiting
    root.mainloop()
    # end of program
    
