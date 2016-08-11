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

  void append(int val){
    node * new_node =  new node(val);
    node * last = this;
    while(last->next != NULL) last = last->next;
    last->next = new_node;
    new_node->before=last;
  }

  void print_all(){
    node *k=this;
    while (k != NULL) {
      cout<< k->value<<endl;
      k=k->next;
    }
  }


};


int main(){

  node *n = new node(2);
  n->append(4);
  n->append(3);
  n->append(6);
  n->print_all();
  return 0;
}
