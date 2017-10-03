#include <iostream>
#include <string>
#include <map>


class Solution {
public:
    bool judgeCircle(std::string moves) {
        std::map < char, int > dict;
        for (int i = 0; i < moves.size(); ++i)
        {
        	/* code */
        	dict[moves[i]] += 1;
        }
        if ((dict['D'] == dict['U']) && (dict['L'] == dict['R']))
        {
        	return true;
        }
        else {
        	return false;
        }
    }
};

int main() {
	Solution p;
	std::string s("LL");
	std::cout << p.judgeCircle(s);
	return 0;
}