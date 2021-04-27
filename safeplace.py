import sys

def main(argv):
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

    while counter < len(argv):
        arg = argv[counter];
        for val in arg:
            final_print = final_print + dict[val];

        if counter != len(argv) - 1:
            final_print = final_print + ",";

        counter += 1;

    print(final_print)

if __name__ == "__main__":
   main(sys.argv[1:])
