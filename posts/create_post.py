"""
adds a new entry to blog.html
"""
def add_entry_to_list() :
  md     = open('4.md')
  header = md.readline()[4:-1] #slice omits the beginning hashes and trailing \n
  md.close()
  html = open('../blog.html')
  lines = html.readlines()
  for i in range(0,len(lines)) :
    if lines[i] == '</ol>\n' :
      lines[i -1] = ' <li><a href="posts/4.html">'+header+'</a></li>\n'
      print('made it')
      break
  html.close()
  html = open('../blog.html','wt')
  html.writelines(lines)
  html.close()

add_entry_to_list()
