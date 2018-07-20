#!/usr/bin/python3
from markdown import markdown
from sys      import argv

"""
adds a new entry to blog.html
returns True if blog.html was written
"""
def add_entry_to_list(post_num) :
  post_num = str(post_num)
  md     = open(post_num+'.md')
  header = md.readline()[4:-1] #slice omits the beginning hashes and trailing \n
  md.close()
  html = open('../blog.html')
  lines = html.readlines()
  html.close()
  new_element = ' <li><a href="posts/'+post_num+'.html">'+header+'</a></li>\n'
  for i in range(0,len(lines)) :
    if lines[i] == new_element:
      print('list item already exists')
      return
    if lines[i] == '</ol>\n' :
      lines.insert(i,new_element)
      break
  html = open('../blog.html','wt')
  html.writelines(lines)
  html.close()
  return True

def add_entry(post_num) :
  try :
    md = open(post_num+'.md')
  except :
    print(post_num+'.md'+' does not exist!')
    return
  html = markdown(md.read())
  md.close()
  post = open(post_num+'.html','wt')
  post.write(html)
  post.close()
  add_entry_to_list(post_num)

if __name__ == '__main__' :
  post_num = argv[1]
  add_entry(post_num)
