// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.My8110
{
    public static class ExampleFunc
    {
        public static Task InvokeAsync(ExampleFuncArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync("my8110::exampleFunc", args ?? new ExampleFuncArgs(), options.WithDefaults());
    }


    public sealed class ExampleFuncArgs : global::Pulumi.InvokeArgs
    {
        [Input("enums")]
        private List<Union<string, Pulumi.My8110.MyEnum>>? _enums;
        public List<Union<string, Pulumi.My8110.MyEnum>> Enums
        {
            get => _enums ?? (_enums = new List<Union<string, Pulumi.My8110.MyEnum>>());
            set => _enums = value;
        }

        public ExampleFuncArgs()
        {
        }
        public static new ExampleFuncArgs Empty => new ExampleFuncArgs();
    }
}
