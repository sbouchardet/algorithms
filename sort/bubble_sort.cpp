#include <iostream>

using namespace std;

class node{
public:
  float value;
  float index;
  node *next;
  node *before;
  node(float val){
    this->value=val;
    this->next=NULL;
    this->before=NULL;
  }
};


int main(){
  node na(20);
  na.value = 10;
  node *n = new node(2);
  n->value=5;
  cout<< n->value<<" e "<<na.value<<endl;
  return 0;
}
