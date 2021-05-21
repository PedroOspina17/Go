
Write a method that takes a string and trending words (top trending words from Twitter etc.) 
and return a response for how many times each trending word is repeated in the given string. 

Ex: "Dog is chasing the cat and cat climbed on a tree", ["cat","small","dog","running"] 
expected response: "[cat - 2, small - 0, dog - 1, running - 0]"


public void countWords(string text,dictionary<string,int> trendingWords ){

    if(text == null ){
        throw new System.BadArgumentException("The text cannot be null");
        }

    if(trendingWords == null ){
        throw new System.BadArgumentException("The trending cannot be null");
        }
        
        &&  != null){
        var words = text.Split(' ')

        foreach (var currentWord in words)
        {
            if(trendingWords.key.find(currentWord)){
                trendingWords[currentWord] += 1;
            }        
            
        }
    
    }
    return trendingWords;
}


