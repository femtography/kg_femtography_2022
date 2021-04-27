import sys

# Main function that takes in arguements from command line and returns phonetic conversion.
def main(argv):
    # Creating three variables. A string for stdout. A counter to trigger a stop in th while loop.
    # As well as the hashmap to convert numerical values to their phonetic counterpart.
    final_print = "";
    counter = 0;
    dict = {
        "0" : "Zero",
        "1" : "One",
        "2" : "Two",
        "3" : "Three",
        "4" : "Four",
        "5" : "Five",
        "6" : "Six",
        "7" : "Seven",
        "8" : "Eight",
        "9" : "Nine"
    }

    # While loop to go through arguements provided from command line. It will use hash table to convert each value
    # provided into the spelling of the number and then add it to final_print string.
    while counter < len(argv):
        arg = argv[counter];
        for val in arg:
            final_print = final_print + dict[val];

        # Adds comma after each value in args but only if it is not the last item in list.
        if counter != len(argv) - 1:
            final_print = final_print + ",";

        counter += 1;

    # Returning expected value.
    print(final_print)

if __name__ == "__main__":
   main(sys.argv[1:])
