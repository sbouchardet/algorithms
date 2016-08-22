#include "structure.cpp"

#include <iostream>

using namespace std;

class insetion_sort{

private:
  static void change(node * n1, node *n2){
    float *a = n1->get_info();
    n1->set_info(n2->get_info());
    n2->set_info(a);
  }
public:
  static node * sort(node * r){
    node * i = r;
    node * j = NULL;
    node * nx = NULL;
    while (i->next != NULL){
      j=r;
      while(j->get_value()<i->get_value())j=j->next;
      nx = i->next;
      r->remove(i);
      if(j==r) r = j->add_before(i);
      j->add_before(i);
      i=nx;
    }
    return r;

  }



};

int main(){
  node * n = new node(-1);
  n->append(1);
  n->append(5);
  n->append(4);
  n->append(2);
  n->append(5);

  n->print_all();
cout<<"\nSorted:"<<endl;
n = insetion_sort::sort(n);
n->print_all();


  return 0;
}
