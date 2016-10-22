import sys
import glob

types = ['uint%d' % s for s in (8, 16, 32, 64)]
types += ['uint']
types += ['int%d' % s for s in (8, 16, 32, 64)]
types += ['int']
types += ['float32', 'float64']


print " ".join(types)

