#include <vector>
// Employee info
class Employee {
public:
    // It's the unique ID of each node.
    // unique id of this employee
    int id;
    // the importance value of this employee
    int importance;
    // the id of direct subordinates
    vector<int> subordinates;
};

class Solution {
public:
    int getImportance(vector<Employee*> employees, int id) {
    	int importance_sum;

    	importance_sum += employees[id-1]->importance;
    	for (int i = 0; i < employees[id-1]->subordinates.size(); i++) {

    		importance_sum += getImportance(employees, employees[id-1]->subordinates[i]);
    	}
        
        return importance_sum;
    }
};