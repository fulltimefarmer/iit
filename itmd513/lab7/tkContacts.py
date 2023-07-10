import os
from tkinter import *
from tkinter import messagebox
from contacts import *

# Jun Zhou
# July 10th, 2023
# Lab7
# ITMD513

def selection():
    print ('At %s of %d' % (select.curselection(), len(contactlist)))
    return int(select.curselection()[0])

def addContact():
    contactlist.append([nameVar.get(), phoneVar.get()])
    setList()
    messagebox.showinfo("Success", "Contacts added successfully!")

def updateContact():
    contactlist[selection()]=[nameVar.get(), phoneVar.get()]
    setList()
    messagebox.showinfo("Success", "Contacts Updated successfully!")

def deleteContact():
    result = messagebox.askquestion('Confirmation', 'Are you sure you want to delete this contact?')
    if result == 'yes':
        del contactlist[selection()]
    setList()

def loadContact():
    name, phone = contactlist[selection()]
    nameVar.set(name)
    phoneVar.set(phone)

def saveContacts():
    temp_file = open('temp.py', 'w')
    temp_file.write('contactlist = [\n')
    for index, contact in enumerate(contactlist):
        formatted = ', '.join(f"'{element}'" for element in contact)
        temp_file.write(f'  [{formatted}]')
        if index != len(contactlist) - 1:
            temp_file.write(',')
        temp_file.write('\n')
    temp_file.write(']')
    os.remove('contacts.py')
    os.rename('temp.py', 'contacts.py')

def exitProgram():
    result = messagebox.askquestion('Exit', 'Are you want to exit?')
    if result == 'yes':
        os._exit(1)

def buildFrame():
    global nameVar, phoneVar, select
    root = Tk()
    root.title('My Contact List')  # Set title
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
    btn2 = Button(frame1,text='Update',command=updateContact)
    btn3 = Button(frame1,text='Delete',command=deleteContact)
    btn4 = Button(frame1,text=' Load ',command=loadContact)
    btn5 = Button(frame1, text=' Save ', command=saveContacts)
    btn1.pack(side=LEFT)
    btn2.pack(side=LEFT)
    btn3.pack(side=LEFT)
    btn4.pack(side=LEFT)
    btn5.pack(side=LEFT)
    frame1 = Frame(root) # allow for selection of names
    frame1.pack()
    scroll = Scrollbar(frame1, orient=VERTICAL)
    select = Listbox(frame1, yscrollcommand=scroll.set, height=7)
    scroll.config (command=select.yview)
    scroll.pack(side=RIGHT, fill=Y)
    select.pack(side=LEFT,  fill=BOTH)
    frame1 = Frame(root) # add exit button
    frame1.pack()
    btn6 = Button(frame1, text=' Exit ', command=exitProgram)
    btn6.pack()
    return root

def setList():
    contactlist.sort()
    select.delete(0, END)
    for name,phone in contactlist:
        select.insert(END, name)

if __name__ == "__main__":
    root = buildFrame()
    setList()
    root.mainloop()
