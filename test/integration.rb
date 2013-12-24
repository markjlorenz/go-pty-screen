#! /usr/bin/env ruby

require 'rspec'
require 'pty'
require 'timeout'
require 'shellwords'

module EventualIO
  READ_SIZE = 1024
  module_function

  @@debug = false
  def debug=(val); @@debug = val; end
  def debug?; @@debug == true; end

  def match_in_time?(io, regex, buffer, timeout=6)
    timeout ||= 6
    happens_in_time(timeout) do
      begin
        fill_buffer(io, buffer) { return false }
      end until buffer.match(regex)
    end
  end

  def happens_in_time(timeout=3)
    Timeout::timeout(timeout) {
      yield
      true
    }
  rescue  Timeout::Error
    false
  end

  # read from io until the block returns true
  def sleep_until(io, timeout=6)
    buffer = ''
    Timeout::timeout(timeout) {
      begin
        fill_buffer(io, buffer) { |e| raise e }
      end until yield(buffer.clone)
    }
  end

  def fill_buffer(io, buffer)
    buffer <<  io.readpartial(READ_SIZE)
    STDOUT.write buffer if debug?
  rescue EOFError => e
    yield(e)
  end
end

RSpec::Matchers.define :eventually_match do |*attrs|
  regex   = attrs[0]
  timeout = attrs[1] || 6
  screen  = ""

  match do |actual|
    screen.clear
    EventualIO.send(:match_in_time?, actual, regex, screen, timeout)
  end

  failure_message_for_should { |actual| Shellwords.escape(screen) }
end

describe "basic screen sharing" do

  describe "`new` creates a new application" do
    before(:all) do
      @server_stdout, @server_stdin, @s_pid = PTY.spawn 'go run go-pty-server.go'
      EventualIO.sleep_until(@server_stdout) { |screen| screen.match /Available.+PTYs/ }
    end

    after(:all) do
      Process.kill("KILL", @s_pid)
    end

    it "registers the first application" do
      @server_stdin << 'new b1 bash 20 80'
      expect(@server_stdout).to eventually_match(/b1.+bash.+\d{4,}.+\d{4,}/)
    end

    it "registers the second application" do
      @server_stdin << 'new b2 bash 20 80'
      expect(@server_stdout).to eventually_match(/b2.+bash.+\d{4,}.+\d{4,}/)
    end

    it "registers the third application" do
      @server_stdin << 'new b3 bash 20 80'
      expect(@server_stdout).to eventually_match(/b3.+bash.+\d{4,}.+\d{4,}/)
    end

    context "a client connects" do
      before(:all) do
        @client1_stdout, @client1_stdin, @c1_pid = PTY.spawn 'go run go-pty-client.go'
        @client2_stdout, @client2_stdin, @c2_pid = PTY.spawn 'go run go-pty-client.go'
        @client3_stdout, @client3_stdin, @c3_pid = PTY.spawn 'go run go-pty-client.go'
      end

      it "can see both apps" do
        expect(@client1_stdout).to eventually_match(/b1.+bash.+b2.+bash/, 8)
        expect(@client2_stdout).to eventually_match(/b1.+bash.+b2.+bash/, 8)
        expect(@client3_stdout).to eventually_match(/b1.+bash.+b2.+bash/, 8)
      end

      it "can connect to b2" do
        @client1_stdin << 'j'
        @client2_stdin << 'j'
        @client3_stdin << 'j'

        @client1_stdin << 'echo $LINES'
        expect(@client1_stdout).to eventually_match(/\s20\s/, 3)
        expect(@client2_stdout).to eventually_match(/\s20\s/, 3)
      end

      it "doesn't crash if client one closes his terminal" do
        Process.kill("KILL", @c1_pid)
        expect{ Process.getpgid @c2_pid }.to_not raise_error
      end

      it "kicks out client 3 if client 2 quits the application" do
        Thread.new do
          process_status = Process.wait2(@c3_pid)[1]
          expect(process_status).to be_success
        end
        @client2_stdin << 'exit'
      end

      it "removes the closed app from the supervisor" do
        expect(@server_stdout).to eventually_match(/b1(?!.+b2).+b3/)
      end

      it "doesn't show clients the removed app" do
        @client4_stdout, @client4_stdin, @c4_pid = PTY.spawn 'go run go-pty-client.go'
        expect(@client4_stdout).to eventually_match(/b1(?!.+b2).+b3/)
      end

    end

    context "starting multiple apps from http" do
      it "starts two new apps" do
        `nc localhost 2000 < test/create-test-3.http`
        expect(@server_stdout).to eventually_match(/b4.+bash.+\d{4,}.+\d{4,}/)
        expect(@server_stdout).to eventually_match(/vim-2.+vim.+\d{4,}.+\d{4,}/)
      end
    end

  end

  describe "loading the .rc file" do
    before(:all) do
      @server_stdout, @server_stdin, @s_pid =
        PTY.spawn 'go run go-pty-server.go --config-file=test/create-test-3.http'
    end

    after(:all) do
      Process.kill("KILL", @s_pid)
    end

    it "registers the rc'd applications" do
      expect(@server_stdout).to eventually_match(/vim-2.+vim.+\d{4,}.+\d{4,}/)
      expect(@server_stdout).to eventually_match(/b4.+bash.+\d{4,}.+\d{4,}/)
    end
  end

end
