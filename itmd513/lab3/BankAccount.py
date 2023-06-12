# jzhou 2023-06-11  ITMD-513
print("%s" % "Bank Account Program")

pin = "1234"            # Set the correct PIN number
success = False         # Set the initial success flag
MAX_ATTEMPTS = 3        # Set the maximum number of attempts
balance = 0.00          # Set the initial balance
interest_rate = 0.00    # Set the initial interest rate

for i in range(MAX_ATTEMPTS):
    pin_input = input("Enter your PIN number ({} tries remaining): ".format(3 - i))
    if pin_input == pin:
        print("You are success login!")
        success = True
        break
    else:
        if i < MAX_ATTEMPTS - 1:
            print("Incorrect PIN number. Please try again.")
        else:
            print("You have exceeded the maximum number of attempts.")
            break
        
if success:
    balance = float(input("Please enter an initial account balance:"))
    interest_rate = float(input("Include annual interest rate (as a decimal):"))
    monthly_interest_rate = interest_rate / 12

    print("Month \t\t Interest Amount \t Balance")
    for month in range(1, 13):
        interest = balance * monthly_interest_rate
        balance = balance + interest
        print("%d \t\t $%.2f \t\t\t $%.2f" % (month, interest, balance))
