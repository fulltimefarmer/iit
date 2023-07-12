import os
from tkinter import *
from tkinter import messagebox
from myDatabase import *

# Jun Zhou
# July 12th, 2023
# Lab8
# ITMD513

stack = []

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
    name, phone = nameVar.get(), phoneVar.get()
    try:
        print('contact number before add operation: %d' % len(contacts))
        createContact(name, phone)
    except:
        messagebox.showwarning("Warning", "Please enter a value in both the fields.")
    else:
        print("New contact whit Name: %s and Phone: %s added." % (nameVar.get(), phoneVar.get()))
        setList()
        print('contact number after add operation: %d' % len(contacts))
        print('All contacts after add opertion:')
        for record in contacts:
            print(record)


def editContact():
    try:
        print('contact number before update operation: %d' % len(contacts))
        contact = findById(currentId)
        updateContact(currentId, nameVar.get(), phoneVar.get())
        stack.append(('EDIT', currentId, contact[1], contact[2], nameVar.get(), phoneVar.get()))
    except e:
        print(e)
        messagebox.showwarning("Warning", "Please enter valid value in both the fields.")
    else:
        print("Success update record from Name: %s and Phone: %s to Name: %s and Phone: %s." % (contact[1], contact[2], nameVar.get(), phoneVar.get()))
        setList()
        print('contact number after update operation: %d' % len(contacts))
        print('All contacts after edit opertion:')
        for record in contacts:
            print(record)


def removeContact():
    _, name, phone = selection()
    result = messagebox.askquestion('Confirmation', ('Are you sure you want to delete %s?' % name))
    if result == 'yes':
        print('contact number before delete operation: %d' % len(contacts))
        deleteContact(currentId)
        stack.append(('DELETE', currentId, name, phone))
        print("Success delete record Name: %s and Phone: %s." % (name, phone))
        setList()
        print('contact number after delete operation: %d' % len(contacts))
        print('All contacts after delete opertion:')
        for record in contacts:
            print(record)


def rollback():
    if len(stack) > 0:
        operation = stack.pop()
        if operation[0] == 'EDIT':
            result = messagebox.askquestion('Confirmation', ('Are you sure you want to revert edit opertion for %s?' % operation[4]))
            if result == 'yes':
                updateContact(operation[1], operation[2], operation[3])
                print("Reverted EDIT operation: Updated record from Name: %s and Phone: %s to Name: %s and Phone: %s." % (operation[4], operation[5], operation[2], operation[3]))
                setList()
        elif operation[0] == 'DELETE':
            result = messagebox.askquestion('Confirmation', ('Are you sure you want to revert delete opertion for %s?' % operation[2]))
            if result == 'yes':
                createContact(operation[2], operation[3])
                print("Reverted DELETE operation: Added contact with Name: %s and Phone: %s." % (operation[2], operation[3]))
                setList()
        print('All contacts after undo opertion:')
        for record in contacts:
            print(record)
    else:
        messagebox.showwarning("Warning", "No opertion can revert!")


def loadContact():
    _, name, phone = selection()
    nameVar.set(name)
    phoneVar.set(phone)
    opertion = None


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
    btn4 = Button(frame1,text='Rollback',command=rollback)
    btn5 = Button(frame1,text=' Load ',command=loadContact)
    btn1.pack(side=LEFT)
    btn2.pack(side=LEFT)
    btn3.pack(side=LEFT)
    btn4.pack(side=LEFT)
    btn5.pack(side=LEFT)
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


def setList():
    global contacts
    contacts = readContact()
    contacts.sort(key = lambda x : x[1])
    select.delete(0, END)
    for _, name, _ in contacts:
        select.insert(END, name)


def findById(id):
    for element in contacts:
        if element[0] == id:
            return element


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
    
