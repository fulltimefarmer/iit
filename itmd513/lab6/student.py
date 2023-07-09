# Jun Zhou
# July 04th, 2023
# Lab6
# ITMD513

class Student:

    Scores = {}

    # initializing the constructor method

    def __init__(self, name, grade):
        self.name = name
        self.grade = grade

    def getScores(self):

        answer_key = []
        # read into answer_key list, the answer key from file
        answer_key = [line.strip() for line in open("answers.txt", 'r')]

        student_answers = []
        # read into student_answers list, student answers from file
        student_answers = [line.strip().split(',') for line in open("data.txt", 'r')]
        total_score = 100

        # place additional code statements here for the above function

        #---start your loop processing logic here---#
        for line in student_answers:
            if line[0] == self.getName():
                answers = line[1:]
                for i in range(len(answer_key)):
                    if answer_key[i] != answers[i]:
                        total_score -= 10  # Deduct 10 points for each incorrect answer
        #---end your loop processing logic here---#

        #---continue the class definition#

        Student.Scores[self.getName()] = [self.getGrade(),total_score]

    def getName(self):
        return self.name

    def getGrade(self):
        return self.grade

    @staticmethod
    def sortDict():
        return sorted(Student.Scores.items())

#---end the class definition#

student_objs = [
    Student('Sammy Student', 65),
    Student('Betty sanchez', 45),
    Student('Alice brown', 100),
    Student('tom Schulz', 50),
    Student('Max Zhou', 90)
]

for index in range(len(student_objs)):
    student_objs[index].getScores()

sortList = Student.sortDict()

total_students = 0
total_average_score = 0
total_score_range = 0

for k, v in sortList:
    total_students += 1
    old_score = v[0]
    new_score = v[1]
    average = (v[0] + v[1]) / 2
    total_average_score += average
    score_range = abs(v[0] - v[1])
    total_score_range += score_range
    print(k, "has Old score: ", old_score)
    print(k, "has New score: ", new_score)
    print("Average Grade: ", average)
    print("Grade Range: ", score_range)
    print("\n")

overall_average_grade = total_average_score / total_students
overall_average_grade_range = total_score_range / total_students
print("Overall Average Grade: ", overall_average_grade)
print("Overall Average Grade Range: ", overall_average_grade_range)
