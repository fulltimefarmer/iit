unitSystem = input("Enter E for English units or M for metric units: ")

if unitSystem == "E":
    weight = float(input("Enter weight in pounds: "))
    height = float(input("Enter height in inches: "))
    bmi = (weight * 703) / (height ** 2)
else:
    weight = float(input("Enter weight in kilograms: "))
    height = float(input("Enter height in meters: "))
    bmi = weight / (height ** 2)

print("%s = %.2f" % ("BMI", bmi))

if bmi < 18.5:
    print("The person is underweight")
elif bmi >= 18.5 and bmi <= 24.9:
    print("The person is normal")
elif bmi >= 25 and bmi <= 29.9:
    print("The person is overweight")
else:
    print("The person is obese")
    
#print("Have a nice day!")