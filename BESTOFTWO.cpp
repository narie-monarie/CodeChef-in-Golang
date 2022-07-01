#include <iostream>
using namespace std;

int main() {
	int a, b, c;
	cin >> c;
	while (c--) {
		cin >> a >> b;
		if (a > b) {
			cout << a << endl;

		} else {
			cout << b << endl;
		}
	}
	return 0;
}

