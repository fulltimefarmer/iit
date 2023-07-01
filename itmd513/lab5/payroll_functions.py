import os 

def calculate_gross_pay(rate=None, hours=None):
    if rate is None:
        rate = float(input("Enter the hourly rate: "))
    if hours is None:
        hours = float(input("Enter the number of hours worked: "))

    gross_pay = 0
    if hours > 40:
        overtime_hours = hours - 40
        regular_pay = rate * 40
        overtime_pay = (rate * 1.5) * overtime_hours
        gross_pay = regular_pay + overtime_pay
    else:
        gross_pay = rate * hours

    return gross_pay


def print_all(file_name=None):
    if file_name is None:
        file_name = input("Enter the payroll file name: ")

    try:
        emp_file = open(file_name, "r")
        for line in emp_file :
            emp_info = line.split(" ")
            grossPay = calculate_gross_pay(float(emp_info[2]), float(emp_info[3]))
            print("%s %s %s %s $%4.2f" % (emp_info[0], emp_info[1], emp_info[2], emp_info[3].strip(), grossPay))
        emp_file.close()
    except FileNotFoundError:
        print("Error: Payroll file not found!")
    except ValueError:
        print("Error: Invalid payroll file format!")  


def print_emp(file_name=None):
    if file_name is None:
        file_name = input("Enter the payroll file name: ")

    try:
        emp_file = open(file_name, "r")
        inputFirstName = input("Enter the employee's first name: ")
        inputLastName = input("Enter the employee's last name: ")

        line = emp_file.readline()
        for line in emp_file :
            emp_info = line.split(" ")
            if (inputFirstName == emp_info[0] and inputLastName == emp_info[1]):
                grossPay = calculate_gross_pay(float(emp_info[2]), float(emp_info[3]))
                print("%s %s %s %s $%4.2f" % (emp_info[0], emp_info[1], emp_info[2], emp_info[3].strip(), grossPay))
                break
        emp_file.close()
    except FileNotFoundError:
        print("Error: Payroll file not found!")
    except ValueError:
        print("Error: Invalid payroll file format!")  


def add_emp(file_name=None):
    if file_name is None:
        file_name = input("Enter the payroll file name: ")

    try:
        emp_file = open(file_name, "r")
        temp_file = open("temp.txt", "a")
        inputFirstName = input("Enter the employee's first name: ")
        if inputFirstName == "":
            inputFirstName = "default"
        inputLastName = input("Enter the employee's last name: ")
        if inputLastName == "":
            inputLastName = "default"
        inputRate = input("Enter the employee's rate of pay: ")
        inputHours = input("Enter the employee's hours worked: ")

        employee_exists = False

        for line in emp_file:
            emp_info = line.split(" ")
            if (inputFirstName == emp_info[0] and inputLastName == emp_info[1]):
                employee_exists = True
                break
        
        if employee_exists:
            print("Employee already exists in the payroll!")
            return
        else:
            emp_file.seek(0)
            for line in emp_file:
                temp_file.write(line)
            temp_file.write("\n")
            temp_file.write("%s %s %s %s" % (inputFirstName, inputLastName, inputRate, inputHours))
            print("Employee added successfully!")

        emp_file.close()
        temp_file.close()
        os.remove(file_name)
        os.rename("temp.txt", file_name)
    except FileNotFoundError:
        print("Error: Payroll file not found!")
    except ValueError:
        print("Error: Invalid payroll file format!") 


def delete_emp(file_name=None):
    if file_name is None:
        file_name = input("Enter the payroll file name: ")

    try:
        emp_file = open(file_name, "r")
        temp_file = open("temp.txt", "w")
        inputFirstName = input("Enter the employee's first name: ")
        inputLastName = input("Enter the employee's last name: ")

        employee_exists = False

        for line in emp_file :
            emp_info = line.split(" ")
            if (inputFirstName == emp_info[0] and inputLastName == emp_info[1]):
                employee_exists = True
                continue
            else:
                temp_file.write(line)
        
        emp_file.close()
        temp_file.close()
        os.remove(file_name)
        os.rename("temp.txt", file_name)
        if employee_exists:
            print("Employee deteled successfully!")
        else:
            print("Employee %s %s not exist!" % (inputFirstName, inputLastName))
    except FileNotFoundError:
        print("Error: Payroll file not found!")
    except ValueError:
        print("Error: Invalid payroll file format!") 
 

def modify_emp(file_name=None):
    if file_name is None:
        file_name = input("Enter the payroll file name: ")

    try:
        emp_file = open(file_name, "r")
        temp_file = open("temp.txt", "w")
        inputFirstName = input("Enter the employee's first name: ")
        inputLastName = input("Enter the employee's last name: ")

        found = False

        emp_file.seek(0)
        for line in emp_file:
            emp_info = line.split(" ")
            if (inputFirstName == emp_info[0] and inputLastName == emp_info[1]):
                found = True
                newRate = input("Enter the new hourly rate: ")
                if newRate == "":
                    newRate = emp_info[2]
                newHours = input("Enter the new number of hours worked: ")
                if newHours == "":
                    newHours = emp_info[3]
                temp_file.write("%s %s %s %s\n" % (inputFirstName, inputLastName, newRate, newHours))
            else:
                temp_file.write(line)

        if found:
            print("Employee record updated successfully!")
        else:
            print("Employee %s %s not found in the employees." % (inputFirstName, inputLastName))

        emp_file.close()
        temp_file.close()
        os.remove(file_name)
        os.rename("temp.txt", file_name)
    except FileNotFoundError:
        print("Error: Payroll file not found!")
    except ValueError:
        print("Error: Invalid payroll file format!") 
