import os 
	
def menu():
    pstr = "Welcome to the Payroll Program\n"
    pstr += "Choose from the following payroll choices\n"
    pstr += "(1) A gross PR payroll report for all employees\n"
    pstr += "(2) A gross PR payroll report for a single employee by name\n"
    pstr += "(3) Add a new employee\n"
    pstr += "(4) Delete an exist employee\n"
    pstr += "(5) Modify an exist employee\n"
    pstr += "(6) Quit Program"
    print (pstr)
   		
def printall():
    empFile = open("employees.txt", "r")
    for line in empFile :
        empInfo = line.split(" ")
        grossPay = float(empInfo[2]) * float(empInfo[3])
        print("%s %s %s %s $%4.2f" % (empInfo[0], empInfo[1], empInfo[2], empInfo[3], grossPay))
    empFile.close()
    
def printEmp():
    empFile = open("employees.txt", "r")
    inputStr = input("Enter in a full name to search employee ")
    line = empFile.readline()
    for line in empFile :
        empInfo = line.split(" ")
        grossPay = float(empInfo[2]) * float(empInfo[3])
        if (inputStr == empInfo[0] + " " + empInfo[1]):
            print("%s %s %s %s $%4.2f" % (empInfo[0], empInfo[1], empInfo[2], empInfo[3], grossPay))
            break
    empFile.close()
    
def addEmp():
    inputStr = input("Enter in a name to search employee ")

def deleteEmp():
    empFile = open("employees.txt", "r")
    temFile = open("temp.txt", "w")
    inputStr = input("Enter in a full name to delete the employee ")
    for line in empFile :
        empInfo = line.split(" ")
        if (inputStr != empInfo[0] + " " + empInfo[1]):
            temFile.write(line)
    os.remove("employees.txt")
    os.rename("temp.txt", "employees.txt")
 
def modifyEmp():
    inputStr = input("Enter in a name to search employee ")

def exitApp():
    os._exit(0)

def main() :
    menu()
    choice = int(input("Enter Menu Choice Now! "))
    if (choice == 1) :
        printall()
    elif (choice == 2) :
        printEmp()
    elif (choice == 3) :
        addEmp()
    elif (choice == 4) :
        deleteEmp()
    elif (choice == 5) :
        modifyEmp()
    elif (choice == 6) :
        exitApp()
    else :
        print("Invalid Choice")
        exitApp()

main()

