#!/usr/bin/python
import pygame
from pygame.locals import *
import math
import random
import time
import os,sys
import argparse

#colors
ALPHA = (0,0,0,0)
RED = (255,0,0)
BLUE = (0,0,255)
GREEN = (0,255,0)
BLACK = (0,0,0)
WHITE = (255,255,255)
YELLOW = (0,255,255)
#create command line parser
parser = argparse.ArgumentParser(description='Create a tree')
parser.add_argument('depth',type=int)
parser.add_argument('length',type=int)
parser.add_argument('angle',type=int)
parser.add_argument('branches',type=int)
parser.add_argument('initialX',type=int)
parser.add_argument('initialY',type=int)
#setup pygame
os.environ["SDL_VIDEODRIVER"]="dummy" #used so it can run on headless server
pygame.init()
pygame.display.init()
screen = pygame.display.set_mode((1,1))
screen_size = (1920,1080)
make_tree = False
tree_surf = pygame.Surface(screen_size).convert_alpha()
tree_surf.fill(ALPHA)
#variables
rand_build = False
font = pygame.font.SysFont('ariel',30)
colors = [RED,GREEN,BLUE]
print_lines={}
variables = {'d':5,'l':50,'a':45,'b':2}
selected = None
number_dict = {}
colored = False
headless = True # if this is being ran on a server with no screen

"""
double linked tree structure with variable children
"""
class Node(object) :
  def __init__(self,x,y,angle,parent) :
    self.x = x
    self.y = y
    self.parent = parent #not used as of now
    self.children = []
    self.angle = angle
"""
creates a surface for pygame to render text in
"""
class Text(object) :

  def __init__(self,text,pos) :
    self.text = text
    if colored :
      self.color = WHITE
    else :
      self.color = BLACK
    self.selected = False
    self.pos=pos
    self.line = None
  def render(self,num) :
    self.line=font.render(self.text+str(num),False,self.color)
    if colored:
      self.color=WHITE
    else :
      self.color =BLACK


for i in range(0,10) :
  number_dict[i+48] = i #map 0-10 to their equalivant keycodes

"""
calculate and draw a polar line
"""
def pol_line(x0,y0,length,angle,color):
  theta = (math.pi / 180.0) * angle
  x1= length*(math.cos(theta))+x0
  y1 = length*(math.sin(theta))+y0
  pygame.draw.line(tree_surf,color,(x0,y0),(x1,y1))
  #screen  .blit(GUI_surf,(0,0))
  if not headless :
    screen.blit(tree_surf,(0,0))
    pygame.display.update()

  return x1,y1

"""
make a binary tree out of Nodes
"""
def grow_tree(parent,depth,branches,length,theta,color) :
  if depth == 0 :#or parent.angle > 60 and parent.angle < 120:
    return
  else :
    if branches % 2 == 1 :
      odd = True
      mid = ((branches - 1) /2) + 1
    else :
      odd = False
      mid = branches/2
    for i in range(1,branches+1) :
      x1,y1 = pol_line(parent.x,parent.y,length,parent.angle,color)
      if not odd: #if there are an even amount of branches
        if i <= mid :
          new_angle=parent.angle+(theta*i)
        else :
          new_angle=parent.angle-(theta*(i-mid))
        if(i>branches) :
          new_angle *= i
      else : #odd number of branches
        if i<mid :
          new_angle = parent.angle - theta*i
        elif i == mid :
          new_angle = parent.angle
        elif i> mid :
          new_angle = parent.angle + (theta*(i-mid))
      new_color = []
      for n in color :
        new_val = 15 + n
        if new_val>= 175 :
          new_val = 175
        new_color.append(new_val)

      new_color = tuple(new_color)
      parent.children.append(Node(x1,y1,new_angle,parent))
      grow_tree(parent.children[-1],depth-1,branches,length,theta,new_color)

"""
runs a gui for creating trees. Only useful if running the program locally
"""
def run_gui():
  show_GUI = True
  screen = pygame.display.set_mode(screen_size,0,32)
  GUI_surf = pygame.Surface((400,200)).convert_alpha()

  if colored :
    BACKGROUND =  BLACK
  else :
    BACKGROUND = WHITE

  screen.fill(BACKGROUND)

  texts = {}
  texts['d']=Text("(d)epth: ",(0,10))
  texts['b']=Text("(b)ranches: ",(0,35))
  texts['l']=Text("(l)ength: ",(200,10))
  texts['a']=Text("(a)ngle: ",(200,35))
  texts['r']=Text("(r)andom: ",(0,60))
  texts['g']=Text("(G)UI: ",(200,60))
  value = ""

  while True:
    mouse_x,mouse_y = pygame.mouse.get_pos()
    if show_GUI: #values: (rendered font, coordinate to be blitted)
      if selected != None :
        texts[selected].color = YELLOW
      texts['d'].render(variables['d'])
      texts['b'].render(variables['b'])
      texts['l'].render(variables['l'])
      texts['a'].render(variables['a'])
      texts['r'].render(str(rand_build))
      texts['g'].render(str(show_GUI))
      GUI_surf.fill(BACKGROUND)

      for val in texts.values() :
        GUI_surf.blit(val.line,val.pos)
      screen.blit(GUI_surf,(0,0))
      screen.blit(tree_surf,(0,0))
      pygame.display.update()

    if make_tree :
      grow_tree(genesis,variables['d'],variables['b'],variables['l'],variables['a'],color)
      make_tree = False

    for event in pygame.event.get():
      if event.type == pygame.QUIT:
        pygame.quit()
        sys.exit()
        break
      if event.type==KEYDOWN :
        if event.key==K_ESCAPE :
          pygame.quit()
          sys.exit()
          break
        if event.key == K_RETURN :
          if value != "" :
            variables[selected]= int(value)
            value = ""

        if event.key in number_dict.keys() :
          value += str(number_dict[event.key])

        if event.key == 100 :
          selected = 'd'
        if event.key == 98 :
          selected = 'b'
        if event.key == 108 :
          selected = 'l'
        if event.key == 97 :
          selected = 'a'

        if event.key == 115 :#s : save image to .png
          print("saving image...")
          pygame.image.save(
            tree_surf,
            "tree_d-"+str(variables['d'])+"_b-"+str(variables['b'])
            +"_l-"+str(variables['l'])+"_a-"+str(variables['a'])+".png")

        if event.key == 114 :#r : toggle random build
          rand_build = not rand_build
        if event.key == K_SPACE :
          tree_surf.fill(BACKGROUND)
          screen.blit(tree_surf,(0,0))
          tree_surf.fill(ALPHA)
          pygame.display.update()
        if event.key == 103 :#g : toggle GUI
          show_GUI = not show_GUI
          if not show_GUI:
            GUI_surf.fill(BACKGROUND)
          screen.blit(GUIsurf,(0,0))
          screen.blit(tree_surf,(0,0))
          pygame.display.update()
      if event.type==6 : #click
        make_tree = True
        genesis = Node(mouse_x,mouse_y,270,None)
        if colored :
          color = random.choice(colors)
        else :
          color = BLACK

        if rand_build :
          variables['d'] = random.randint(2,6) #depth
          variables['b'] = random.randint(2,5) #branches
          variables['l'] = random.randint(50,150) #length
          variables['a'] = random.randint(10,40) #angle
"""
depth = 3
branches = 4
length = 60
angle = 50
x = 500
y = 600
"""
if not headless : run_gui()

else :
  arguments = parser.parse_args()
  depth = arguments.depth
  branches = arguments.branches
  length = arguments.length
  angle = arguments.angle
  x = arguments.initialX
  y = arguments.initialY
  genesis = Node(x,y,270,None)
  grow_tree(genesis,depth,branches,length,angle,BLACK)
  file = "tree_d-"+str(depth)+"_b-"+str(branches)+"_l-"+str(length)+"_a-"+str(angle)+".png"
  pygame.image.save(tree_surf,file)
  sys.stdout.write(file)

