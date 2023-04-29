#include <stdio.h>
#include <unistd.h>

int count_digits(long long num) {
    int count = 0;
    while (num != 0) {
        num /= 10;
        count++;
    }
    return count;
}

void asker(long long *cuil) {
    printf("Enter your cuil (only numbers): \n");
    
    while (1) {
        scanf("%lld", cuil);
    
        if (count_digits(*cuil) == 11) {
            break;
        }
        else {
            printf("Please enter a correct value\n");
        }
    }
}

void checker(long long cuil){ //I've made this function for the hell of it
    printf("Checking");
    
    for (int i = 0; i < 3; i++) {
        printf(".");
        fflush(stdout); // flush the output buffer to ensure the dots are printed immediately
        sleep(1);
    }
    
    int last_two_digits = cuil % 100;
    printf("\nLast two digits: %d\n", last_two_digits);
}

int main() {
    long long cuil;
    asker(&cuil);
    printf("Your cuil is: %lld\n", cuil);
    
    checker(cuil);

    return 0;
}
