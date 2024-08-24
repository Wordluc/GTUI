package Drawing

import "testing"

func TestTypingInTextBlock(t *testing.T) {
    textBlock:=CreateTextBlock(0,0,10,100)
	 textBlock.Type('a')
	 textBlock.Type('2')
	 textBlock.Type('3')
	 textBlock.Type('\n')
	 textBlock.Type('n')
	 expLine0:=[]rune{'a','2','3'}
	 expLine1:=[]rune{'n'}
	 for i:=0;i<2;i++{
		 if textBlock.lines[0].line[i]!=expLine0[i]{
			 t.Errorf("expected %v, got %v",string(expLine0[i]),string(textBlock.lines[0].line[i]))
		 }
	 }
	 for i,char:=range textBlock.lines[0].getFullText(){
		 if i>=len(expLine0){
			 break
		 }
		 if char!=expLine0[i]{
			 t.Errorf("expected %v, got %v",string(expLine0[i]),string(char))
		 }
	 }
	 for i,char:=range textBlock.lines[1].getFullText(){
		 if i>=len(expLine1){
			 break
		 }
		 if char!=expLine1[i]{
			 t.Errorf("expected %v, got %v",string(expLine1[i]),string(char))
		 }
	 }
 }
