int sumCStyle( int* numberArray , int len )
{
  int total = 0;

  for (int i=0; i<len; i++)
	total += numberArray[i];
}

int main()
{
  const int tries = 100000000;

  int numberArray[21] = {49,50,51,49,50,51,49,50,51,49,50,51,49,50,51,49,50,51,49,50,51};

  int total = 0;
  for (int i=0; i<tries; i++)
	total = sumCStyle( numberArray, 21 );

  total += 1;
}

