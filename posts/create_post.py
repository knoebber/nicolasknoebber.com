#!/usr/bin/python3
from markdown import markdown
from sys      import argv

"""
adds a new <li> element to the <ol> in blog.html
this always adds it to the end
"""
def add_entry_to_list(post_num) :
  post_num = str(post_num)
  md     = open(post_num+'.md')
  header = md.readline()[3:-1] #slice omits the beginning hashes and trailing \n
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
  print('new list item created')
  return True

"""
reads from the markdown file (post_num).md and writes (post_num).html
"""
def add_entry(post_num) :
  try :
    md = open(post_num+'.md')
  except :
    print(post_num+'.md'+' does not exist!')
    return
  html = markdown(md.read())
  html = '<LINK REL=StyleSheet HREF="style.css" TYPE="text/css">\n' + html
  md.close()
  post = open(post_num+'.html','wt')
  post.write(html)
  post.close()
  print('post updated')
  add_entry_to_list(post_num)

if __name__ == '__main__' :
  post_num = argv[1]
  add_entry(post_num)
