'''

Author: David Mathis (dmathis@mediatemple.net)
Date: 10/25/15
Purpose: Parse out some metric data on a given formatted text file

'''
import sys, getopt


def main(argv):

    data_file = ''

    myopts, args = getopt.getopt(sys.argv[1:],"i:")

    # unless in proper syntax display the help
    if len(sys.argv) < 2:
        print("Usage: %s -i data.txt " % sys.argv[0])
        sys.exit()
    for o, a in myopts:
        if o == '-i':
            data_file = a
        else:
            print("Usage: %s -i data.txt" % sys.argv[0])
            sys.exit()

   # slurp up data into a instance of DataParser
    my_data = DataParser(data_file)
    nodes = my_data.get_nodes()
    metrics = my_data.get_metrics()

    # loop through the nodes and return their proper data set
    for node in nodes:
        node_index = nodes.index(node)
        print(node + ": " +
              my_data.get_avg(node_index) +
              my_data.get_max(node_index) +
              my_data.get_min(node_index))

class DataParser:


    # sanitize our data file and store it away (no need for close() in a with)
    def __init__(self, data_file):
        self.data = []
        with open(data_file) as datafile:
            for line in datafile:
                self.data.append(line.strip().replace('|', ',').split(','))

    # grab first value of each sub list as a hostname/node
    def get_nodes(self):
        self.nodes = []
        for item in self.data:
            self.nodes.append(item[0])
        return self.nodes

    # strip out the metric data and sanitize non floats
    def get_metrics(self):
        self.metrics = []
        for item in self.data:
            self.metrics.append([0.0 if i == 'None' else float(i) for i in item[4:]])
        return self.metrics

    def get_max(self, node_index):
        return " Max: " + str(sorted(self.metrics[node_index])[-1])

    def get_min(self, node_index):
        return " Min: " + str(sorted(self.metrics[node_index])[0])

    def get_avg(self, node_index):
        return " Average: " + str(round(sum(self.metrics[node_index]) / float(len(self.metrics[node_index])), 1))

if __name__ == "__main__":
    main(sys.argv[1:])
