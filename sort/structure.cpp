#include <iostream>

using namespace std;

class node{

private:

  float value;
  float index;

public:
  node *next;
  node *before;

  node(float val){

    this->value = val;
    this->index = 0;
    this->next = NULL;
    this->before = NULL;

  }


  void append(int val){
    node * new_node =  new node(val);
    node * last = this;
    while(last->next != NULL) last = last->next;
    last->add_next(new_node);
  }

  void add_next(node * new_node ){
    node * last = this;
    new_node->before=last;
    new_node->next=this->next;
    new_node->index = last->index + 1;
    last->next = new_node;

  }
  node * add_before(node * new_node ){
    node * first = this;
    if(this->before == NULL) first=new_node;

    new_node->before=this->before;
    new_node->next=this;
    new_node->index = this->index + 1;
    this->before = new_node;
    return first;


  }

  node * remove (node * node_to_remove){
    if(node_to_remove->before == NULL) return node_to_remove->next;
    node * bf = node_to_remove->before;
    node * nx = node_to_remove->next;
    bf->next = nx;
    nx->before = bf;
    return this;
  }

  void print_all(){
    node *k=this;
    while (k != NULL) {
      cout<< k->value <<'|'<<k->index<<endl;
      k=k->next;
    }
  }

   float* get_info(){
    float *result = new float[2];
    result[0]=this->index;
    result[1]=this->value;
    return result;
  }

  void set_info(float * info){
    this->index=info[0];
    this->value=info[1];
  }

  float get_value(){
    return this->value;
  }

};


// int main(){
//
//   node *n = new node(2);
//   n->add_next(new node(4));
//   n->add_next(new node(5));
//   n->add_next(new node(1));
//   n->add_next(new node(1));
//   n=n->add_before(new node(-2));
//
//   n->print_all();
//   return 0;
// }
