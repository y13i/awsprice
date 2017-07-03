require "formula"

class AWSPrice < Formula
  VERSION = "0.0.1"

  homepage "https://github.com/y13i/awsprice"
  version  VERSION
  url      "https://github.com/y13i/awsprice/releases/download/#{VERSION}/awsprice-darwin_amd64.zip"
  sha256   "1a187c1dd965b8f4b62de09546caf7623017da7cce43072121605a2187616b37"

  head "https://github.com/y13i/awsprice.git", :branch => "master"

  def install
    system "unzip awsprice-darwin_amd64.zip"
    bin.install "awsprice"
  end
end
