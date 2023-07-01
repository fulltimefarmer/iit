import os 
import payroll_functions

# Jun Zhou
# June 26th, 2023
# Lab5
# ITMD513

def menu():
    """
    Displays the payroll menu options.
    """
    pstr = "Welcome to the Payroll Program\n"
    pstr += "Choose from the following payroll choices\n"
    pstr += "(1) A gross PR payroll report for all employees\n"
    pstr += "(2) A gross PR payroll report for a single employee by name\n"
    pstr += "(3) Add a new employee\n"
    pstr += "(4) Delete an exist employee\n"
    pstr += "(5) Modify an exist employee\n"
    pstr += "(6) Quit Program"
    print (pstr)


def process_option(choice, file_name):
    """
    Processes the user's choice and calls the corresponding payroll function.

    Parameters:
        choice (str): The user's choice.
        file_name (str): The name of the payroll file.
    """
    if (choice == "1") :
        payroll_functions.print_all(file_name)
    elif (choice == "2") :
        payroll_functions.print_emp(file_name)
    elif (choice == "3") :
        payroll_functions.add_emp(file_name)
    elif (choice == "4") :
        payroll_functions.delete_emp(file_name)
    elif (choice == "5") :
        payroll_functions.modify_emp(file_name)
    elif (choice == "6") :
        os._exit(0)
    else :
        print("Invalid Choice")
        os._exit(0)


def main():
    file_name = input("Enter the payroll file name: ")
    menu()
    option = input("Enter your choice (1-6): ")
    process_option(option, file_name)


if __name__ == "__main__":
    main()
